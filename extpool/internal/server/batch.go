package server

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"rock-chain/extpool/internal/client"
	pb "rock-chain/extpool/proto"
	"sync"
	"time"

	"github.com/puzpuzpuz/xsync/v3"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

const COMPLETE_BATCH_THRESHOLD = 10000
// const COMPLETE_BATCH_THRESHOLD = 0

const BATCH_TIMEOUT = 100 * time.Millisecond

// txBatcher 用于聚合 txs
// 当 tx 数量达到 COMPLETE_BATCH_THRESHOLD 时 或者 超时时，就会生成一个 batch
type txBatcher struct {
	// 用于接收新的 txs
	newTxsCh chan []*pb.Transaction
	// 用于发送新的 batch
	newBatchCh chan *pb.Batch

	// 当前 batch 的 txs
	txs []*pb.Transaction
}

func newTxBatcher(newTxsCh chan []*pb.Transaction) *txBatcher {
	return &txBatcher{
		newTxsCh:   make(chan []*pb.Transaction),
		newBatchCh: make(chan *pb.Batch, 10), // 生成 10 个 batch 后不再打包新的 batch
		txs:        make([]*pb.Transaction, 0),
	}
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func genBatchID() string {
	// return 20240219_214205_123_qwer
	currentTime := time.Now()
	year := currentTime.Format("2006")
	month := currentTime.Format("01")
	day := currentTime.Format("02")
	hour := currentTime.Format("15")
	minute := currentTime.Format("04")
	second := currentTime.Format("05")
	// millisecond := currentTime.UnixNano() / 1e6
	millisecond := fmt.Sprintf("%03d", currentTime.UnixNano()/1e6%1000)
	randomString := generateRandomString(4)

	result := fmt.Sprintf("%s%s%s_%s%s%s_%s_%s", year, month, day, hour, minute, second, millisecond, randomString)
	fmt.Println(result)
	return result
}

func (tb *txBatcher) newBatch() {
	batch := &pb.Batch{
		Id:  genBatchID(),
		Txs: tb.txs,
	}
	log.Printf("[extpool] newBatch: %v", batch.Id)
	tb.newBatchCh <- batch
	tb.txs = make([]*pb.Transaction, 0)
}

func (tb *txBatcher) Run() {
	// TODO 超时：第一笔交易到来后，如果 BATCH_TIMEOUT 时间内仍然没有收集到足够的交易，就生成一个 batch
	for {
		select {
		case txs := <-tb.newTxsCh:
			// TODO: 如果 txs 的数量超过 COMPLETE_BATCH_THRESHOLD，就直接生成 batch，不要和现有的 txs 合并
			tb.txs = append(tb.txs, txs...)
			if len(tb.txs) >= COMPLETE_BATCH_THRESHOLD {
				tb.newBatch()
			}
		case <-time.After(BATCH_TIMEOUT):
			if len(tb.txs) > 0 {
				tb.newBatch()
			}
		}
	}
}

// BatchFetcher 读取 batch 拉取请求，向其他节点拉取 batch
type BatchFetcher struct {
	clients []*client.Client // 其他节点的 grpc client

	fetchingBatches *xsync.MapOf[string, bool] // 正在拉取的 batchID
	dirData         string                     // 数据文件夹路径

	logger *zerolog.Logger
}

// FetchBatch 不断拉取 batch，直到成功，然后写入文件
func (bf *BatchFetcher) FetchBatch(batchId string) {
	bf.logger.Info().Str("batchId", batchId).Msg("FetchBatch")
	bf.fetchingBatches.Store(batchId, true)

	for {
		fetch_success := false
		for _, c := range bf.clients {

			if resp, err := c.FetchBatch(&pb.FetchBatchRequest{Id: batchId}); err != nil {
				bf.logger.Error().Err(err).Msg("FetchBatch error")
			} else {
				if resp.Status == 0 {
					WriteBatchToFile(bf.dirData, resp.Batch)
					fetch_success = true
					break
				}
			}
		}
		if fetch_success {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err := os.Remove(fmt.Sprintf("%s/fetching/%s", bf.dirData, batchId)); err != nil {
		bf.logger.Error().Err(err).Msg("Remove file error")
	}

	bf.fetchingBatches.Delete(batchId)
}

func (bf *BatchFetcher) Run() {
	for {
		// 遍历 dirData, 文件列表即为待拉取的 batchID。
		// 对于每个 batchID，如果不在 fetchingBatches 中，就拉取
		if files, err := os.ReadDir(fmt.Sprintf("%s/fetching", bf.dirData)); err != nil {
			bf.logger.Error().Err(err).Msg("ReadDir error")
		} else {
			for _, file := range files {
				batchId := file.Name()
				if _, ok := bf.fetchingBatches.Load(batchId); !ok {
					go bf.FetchBatch(batchId)
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

type BatchManagerConfig struct {
	Others []string // 其他节点的 grpc 地址

	dirData string // 数据文件夹路径
	logger  *zerolog.Logger
}

// BatchManager
// 1. 接受交易，聚合交易，生成 batch
// 2. 产生 ack
// 3. 写入文件，发送给其他节点
type BatchManager struct {
	txBatcher    *txBatcher
	batchFetcher *BatchFetcher

	clients    []*client.Client // 其他节点的 grpc client
	prepareNum int              // 传递给其他 prepareNum 个节点后，再写入文件

	pendingBatches *xsync.MapOf[string, *pb.Batch] // 待 ack 的 batch

	batchSentStatus     map[string][]bool // batchID -> otherIndex -> 是否已发送
	batchSentStatusLock sync.RWMutex

	ackBatchIDs map[string]bool // 已 ack 的 batchID, 只要 key 存在，就表示已 ack

	dirData string // 数据文件夹路径
}

func NewBatchManager(cfg *BatchManagerConfig) *BatchManager {
	clients := make([]*client.Client, len(cfg.Others))
	for i, addr := range cfg.Others {
		s := client.NewClient(addr, cfg.logger)
		clients[i] = s
	}

	newTxsCh := make(chan []*pb.Transaction)
	bm := &BatchManager{
		clients: clients,
		// prepareNum 使用 2，以后改
		prepareNum:      2,
		dirData:         cfg.dirData,
		pendingBatches:  xsync.NewMapOf[string, *pb.Batch](),
		ackBatchIDs:     make(map[string]bool),
		batchSentStatus: make(map[string][]bool),
	}

	bm.txBatcher = newTxBatcher(newTxsCh)
	bm.batchFetcher = &BatchFetcher{
		clients:         clients,
		fetchingBatches: xsync.NewMapOf[string, bool](),
		dirData:         cfg.dirData,
		logger:          cfg.logger,
	}

	return bm
}

// 接受交易，不阻塞
// 返回是否接受成功
func (bm *BatchManager) AddTxs(txs []*pb.Transaction) bool {
	// TODO: 如果 txBatcher 已满，就返回 false
	go func() {
		bm.txBatcher.newTxsCh <- txs
	}()
	return true
}

// processPreparedBatch 处理准备好的 batch
func (bm *BatchManager) processPreparedBatch(batch *pb.Batch) error {
	// 1. 写入文件
	if err := WriteBatchToFile(bm.dirData, batch); err != nil {
		return err
	}

	// 2. 删除 pendingBatches
	bm.pendingBatches.Delete(batch.Id)
	return nil
}

func (bm *BatchManager) AddBatch(batch *pb.Batch) {
	bm.pendingBatches.Store(batch.Id, batch)
}

func (bm *BatchManager) AckBatch(id string) {
	if batch, ok := bm.pendingBatches.Load(id); !ok {
		// TODO: 收集 ack
		return
	} else {
		bm.processPreparedBatch(batch)
	}
}

// Run 启动 batch manager
func (bm *BatchManager) Run() {
	go bm.txBatcher.Run()
	go bm.batchFetcher.Run()

	for batch := range bm.txBatcher.newBatchCh {
		// bm.processPreparedBatch(batch)

		bm.batchSentStatusLock.Lock()
		bm.batchSentStatus[batch.Id] = make([]bool, len(bm.clients))
		bm.batchSentStatusLock.Unlock()

		// 2. 发送给其他节点
		for i, c := range bm.clients {
			go func(i int, c *client.Client) {
				_, err := c.SendBatch(&pb.SendBatchRequest{Batch: batch})
				if err != nil {
					log.Printf("SendTxs error: %v", err)
				} else {
					count := 0
					bm.batchSentStatusLock.Lock()
					status, ok := bm.batchSentStatus[batch.Id]
					if ok {
						status[i] = true
					}
					for _, success := range status {
						if success {
							count++
						}
					}
					bm.batchSentStatusLock.Unlock()
					if count == bm.prepareNum {
						go bm.processPreparedBatch(batch)
						for _, c := range bm.clients {
							go func(c *client.Client) {
								_, err := c.AckBatch(&pb.AckBatchRequest{Id: batch.Id})
								if err != nil {
									log.Printf("AckBatch error: %v", err)
								}
							}(c)
						}
					}
				}
			}(i, c)
		}

		bm.pendingBatches.Store(batch.Id, batch)
	}
}

func WriteBatchToFile(dirData string, batch *pb.Batch) error {
	if dirData == "" {
		panic("dirData is empty")
	}

	// TODO: use os.WriteFile

	start := time.Now()
	fileName := fmt.Sprintf("%s/%v", dirData, batch.Id)
	defer func() {
		elapsed := time.Since(start)
		log.Printf("write file to %v, elapsed: %v", fileName, elapsed)
	}()

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("create file error: %v", err)
		return err
	}
	defer file.Close()
	pbBytes, err := proto.Marshal(batch)
	if err != nil {
		log.Printf("proto marshal error: %v", err)
		return err
	}
	_, err = file.Write(pbBytes)
	if err != nil {
		log.Printf("write file error: %v", err)
		return err
	}
	return nil
}

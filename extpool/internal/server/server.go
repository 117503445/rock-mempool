package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"google.golang.org/protobuf/proto"
	pb "rock-chain/extpool/proto"
)

type Config struct {
	DirData string // 数据文件夹路径，为空时使用默认路径 /tmp/extpool/$pid

	Host   string
	Others []string

	Logger *zerolog.Logger
}

type GrpcServer struct {
	pb.UnimplementedMempoolServer

	cfg          *Config
	batchManager *BatchManager
}

func (g *GrpcServer) AddTxs(ctx context.Context, in *pb.AddTxsRequest) (*pb.AddTxsResponse, error) {
	// 超过 COMPLETE_BATCH_THRESHOLD 个 txs，就认为是一个完整的 batch，否则要等待聚合
	g.cfg.Logger.Info().Str("host", g.cfg.Host).Int("txs", len(in.GetTxs())).Msg("AddTxs Received")

	if !g.batchManager.AddTxs(in.GetTxs()) {
		return &pb.AddTxsResponse{Status: 1}, nil
	}

	return &pb.AddTxsResponse{Status: 0}, nil
}

func (g *GrpcServer) SendBatch(ctx context.Context, in *pb.SendBatchRequest) (*pb.AddTxsResponse, error) {
	batch := in.Batch
	g.cfg.Logger.Info().Str("id", batch.Id).Str("host", g.cfg.Host).Msg("SendBatch Received")

	g.batchManager.AddBatch(batch)

	return &pb.AddTxsResponse{Status: 0}, nil
}

func (g *GrpcServer) AckBatch(ctx context.Context, in *pb.AckBatchRequest) (*pb.AddTxsResponse, error) {
	g.cfg.Logger.Info().Str("id", in.Id).Str("host", g.cfg.Host).Msg("AckBatch Received")

	g.batchManager.AckBatch(in.Id)

	return &pb.AddTxsResponse{Status: 0}, nil
}

func (g *GrpcServer) FetchBatch(ctx context.Context, in *pb.FetchBatchRequest) (*pb.FetchBatchResponse, error) {
	g.cfg.Logger.Info().Str("id", in.Id).Str("host", g.cfg.Host).Msg("FetchBatch Received")

	id := in.Id
	if _, err := os.Stat(fmt.Sprintf("%s/%s", g.cfg.DirData, id)); os.IsNotExist(err) {
		return &pb.FetchBatchResponse{Status: 1}, err
	} else {

		batchBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", g.cfg.DirData, id))
		if err != nil {
			return &pb.FetchBatchResponse{Status: 1}, err
		}

		batch := &pb.Batch{}
		if err := proto.Unmarshal(batchBytes, batch); err != nil {
			return &pb.FetchBatchResponse{Status: 1}, err
		}

		return &pb.FetchBatchResponse{Status: 0, Batch: batch}, nil
	}
}

func (g *GrpcServer) Run() {
	lis, err := net.Listen("tcp", g.cfg.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// max size
	opt := grpc.MaxRecvMsgSize(1024 * 1024 * 1024)

	s := grpc.NewServer(opt)
	pb.RegisterMempoolServer(s, g)
	g.cfg.Logger.Info().Str("host", g.cfg.Host).Msg("server listening")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServer(cfg *Config) *GrpcServer {
	if cfg.Logger == nil {
		panic("logger is nil")
	}

	if cfg.DirData == "" {
		log.Printf("pid: %d", os.Getpid())
		pid := os.Getpid()
		cfg.DirData = fmt.Sprintf("/tmp/extpool/%d", pid)
	}

	// 是否存在
	if _, err := os.Stat(cfg.DirData); os.IsNotExist(err) {
		cfg.Logger.Info().Str("dirData", cfg.DirData).Msg("create dirData")
		// 创建目录
		if err := os.MkdirAll(cfg.DirData+"/fetching", 0755); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		panic(fmt.Sprintf("dirData exists: %s", cfg.DirData))
	}

	bm := NewBatchManager(&BatchManagerConfig{
		Others:  cfg.Others,
		dirData: cfg.DirData,
		logger:  cfg.Logger,
	})
	go bm.Run()

	return &GrpcServer{batchManager: bm, cfg: cfg}
}

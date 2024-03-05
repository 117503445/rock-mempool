package client

import (
	"context"
	"log"
	pb "rock-chain/extpool/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pb.MempoolClient
	target string
	logger *zerolog.Logger
}



func NewClient(target string, logger *zerolog.Logger) *Client {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}
	c := pb.NewMempoolClient(conn)

	return &Client{c,  target, logger}
}

func (c *Client) AddTxs(req *pb.AddTxsRequest) (*pb.AddTxsResponse, error) {
	return c.MempoolClient.AddTxs(context.Background(), req)
}

func (c *Client) SendBatch(req *pb.SendBatchRequest) (*pb.AddTxsResponse, error) {
	return c.MempoolClient.SendBatch(context.Background(), req)
}

func (c *Client) AckBatch(req *pb.AckBatchRequest) (*pb.AddTxsResponse, error) {
	return c.MempoolClient.AckBatch(context.Background(), req)
}

func  (c *Client) FetchBatch(req *pb.FetchBatchRequest) (*pb.FetchBatchResponse, error) {
	c.logger.Info().Str("target", c.target).Msg("FetchBatch")
	return c.MempoolClient.FetchBatch(context.Background(), req)
}

func genNTX() *pb.Transaction {
	return &pb.Transaction{To: []byte("000000000000000000000000000000000000100c"),
		TxType:    pb.Transaction_NTX,
		ExtraData: "01d0439600000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005416c696365000000000000000000000000000000000000000000000000000000",
		Version:   []byte("1")}
}

func GenNTXBatchs(batchNum, batchSize int) [][]*pb.Transaction {
	batchs := make([][]*pb.Transaction, batchNum)
	for i := 0; i < batchNum; i++ {
		batchs[i] = make([]*pb.Transaction, batchSize)
		for j := 0; j < batchSize; j++ {
			batchs[i][j] = genNTX()
		}
	}
	return batchs
}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"rock-chain/extpool/internal/client"
	pb "rock-chain/extpool/proto"

	"github.com/rs/zerolog"
)

func main() {
	const batchNum = 20
	const batchSize = 100000
	batchs := client.GenNTXBatchs(batchNum, batchSize)

	fmt.Println("finish init")

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := client.NewClient("localhost:50051", &logger)

	for i := 0; i < batchNum; i++ {
		req := &pb.AddTxsRequest{
			Txs: batchs[i],
		}

		rpcStart := time.Now()
		r, err := client.AddTxs(req)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		rpcElapsed := time.Since(rpcStart)
		log.Printf("rpc took %s, status: %v", rpcElapsed, r.GetStatus())
	}
}

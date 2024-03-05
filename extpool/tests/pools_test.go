package tests

import (
	"fmt"
	"os"
	"rock-chain/extpool/internal/client"
	"rock-chain/extpool/internal/server"
	pb "rock-chain/extpool/proto"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestPoolsPref(t *testing.T) {
	dirData := "./data"
	if err := os.RemoveAll(dirData); err != nil {
		panic(err)
	}
	os.MkdirAll(dirData, 0755)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	hosts := []string{"localhost:60051", "localhost:60052", "localhost:60053", "localhost:60054"}
	servers := make([]*server.GrpcServer, len(hosts))
	for i, host := range hosts {
		dirNodeData := fmt.Sprintf("%s/%d", dirData, i)

		others := make([]string, 0)
		for j, h := range hosts {
			if j != i {
				others = append(others, h)
			}
		}

		servers[i] = server.NewServer(&server.Config{
			DirData: dirNodeData,
			Host:    host,
			Others:  others,
			Logger:  &logger,
		})
		go func(i int) {
			servers[i].Run()
		}(i)
	}

	time.Sleep(time.Millisecond * 1000) // wait for servers to start

	const batchNum = 5
	const batchSize = 10
	batchs := client.GenNTXBatchs(batchNum, batchSize)

	client := client.NewClient(hosts[0], &logger)

	for i := 0; i < batchNum; i++ {
		req := &pb.AddTxsRequest{
			Txs: batchs[i],
		}
		_, err := client.AddTxs(req)
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		time.Sleep(time.Millisecond * 200)
	}

	time.Sleep(time.Millisecond * 1000) // wait for servers to process

	for i := range hosts {
		dirNodeData := fmt.Sprintf("%s/%d", dirData, i)
		if files, err := os.ReadDir(dirNodeData); err != nil {
			t.Fatalf("could not read dir: %v", err)
		} else {
			if len(files) != batchNum+1 {
				t.Fatalf("files num not match: %v", len(files))
			}
		}
	}
}

func TestPoolsFetchingMissing(t *testing.T) {
	dirData := "./data"
	if err := os.RemoveAll(dirData); err != nil {
		panic(err)
	}
	os.MkdirAll(dirData, 0755)

	hosts := []string{"localhost:60051", "localhost:60052", "localhost:60053", "localhost:60054"}
	servers := make([]*server.GrpcServer, len(hosts))
	for i, host := range hosts {
		dirNodeData := fmt.Sprintf("%s/%d", dirData, i)

		others := make([]string, 0)
		for j, h := range hosts {
			if j != i {
				others = append(others, h)
			}
		}

		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		logger := zerolog.New(output).With().Timestamp().Logger().With().Str("host", host).Logger()

		servers[i] = server.NewServer(&server.Config{
			DirData: dirNodeData,
			Host:    host,
			Others:  others,
			Logger:  &logger,
		})
		go func(i int) {
			servers[i].Run()
		}(i)
	}

	time.Sleep(time.Millisecond * 1000) // wait for servers to start

	txs := client.GenNTXBatchs(1, 10)
	batch := &pb.Batch{
		Id:  "test-batch",
		Txs: txs[0],
	}

	for i := 0; i < len(hosts)-1; i++ {
		server.WriteBatchToFile(fmt.Sprintf("%s/%d", dirData, i), batch)
	}

	os.WriteFile(fmt.Sprintf("%s/%d/fetching/%v", dirData, len(hosts)-1, "test-batch"), []byte(""), 0644)

	time.Sleep(time.Millisecond * 1000) // wait for servers to process

	fileNode3Batch := fmt.Sprintf("%s/%d/%v", dirData, len(hosts)-1, "test-batch")
	if _, err := os.Stat(fileNode3Batch); os.IsNotExist(err) {
		t.Fatalf("file not exists: %v", fileNode3Batch)
	}

	// for i := range hosts {
	// 	dirNodeData := fmt.Sprintf("%s/%d", dirData, i)
	// 	if files, err := os.ReadDir(dirNodeData); err != nil {
	// 		t.Fatalf("could not read dir: %v", err)
	// 	} else {
	// 		if len(files) != batchNum + 1 {
	// 			t.Fatalf("files num not match: %v", len(files))
	// 		}
	// 	}
	// }
}

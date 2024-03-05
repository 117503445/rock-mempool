# extpool

Golang 实现的外置 mempool

## dev

构建开发环境

```sh
docker compose up -d
```

使用 VSCode Attach 到开发容器的 /workspace 目录

## 编译 proto 文件

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/rock.proto
```

## 常用开发操作

```sh
# 运行 mempool
go run ./cmd/mempool/mempool.go

# 构建 mempool
go build cmd/mempool/mempool.go

# 构建小尺寸的 mempool
# https://gophercoding.com/reduce-go-binary-size
# https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md
go build -ldflags="-s -w" -gcflags=all=-l cmd/mempool/mempool.go

# 运行压测端
go run ./cmd/sender/sender.go

# 构建压测端
go build ./cmd/sender/sender.go

# 删除交易缓存
rm -r /tmp/extpool

# 运行 4 个 mempool
HOST="localhost:50051" OTHERS="localhost:50052,localhost:50053,localhost:50054" go run ./cmd/mempool/mempool.go
HOST="localhost:50052" OTHERS="localhost:50051,localhost:50053,localhost:50054" go run ./cmd/mempool/mempool.go
HOST="localhost:50053" OTHERS="localhost:50051,localhost:50052,localhost:50054" go run ./cmd/mempool/mempool.go
HOST="localhost:50054" OTHERS="localhost:50051,localhost:50052,localhost:50053" go run ./cmd/mempool/mempool.go
```

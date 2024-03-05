# fastpool

使用 C++ GRPC 实现的高性能 mempool

## dev

构建开发环境

```sh
docker compose up -d
```

使用 VSCode Attach 到开发容器的 /workspace 目录

## 使用

```sh
# 构建二进制 `./build/mempool` 和 `./build/client`
./build.sh

# 运行 4 个 mempool
HOST="localhost:50051" OTHERS="localhost:50052,localhost:50053,localhost:50054" ./mempool
HOST="localhost:50052" OTHERS="localhost:50051,localhost:50053,localhost:50054" ./mempool
HOST="localhost:50053" OTHERS="localhost:50051,localhost:50052,localhost:50054" ./mempool
HOST="localhost:50054" OTHERS="localhost:50051,localhost:50052,localhost:50053" ./mempool
```

mempool 默认会将 Batch 保存在 /tmp/$PID 目录下

`client` 是压测端，向 `localhost:50051` 发送交易

# rock-mempool

外置交易池

`extpool` 是使用 Go 实现的外置交易池

`fastpool` 是使用 C++ 实现的外置交易池

## Insight

让交易体本身只发送一次，以后的网络消息中只发送交易列表的 ID。从而实现网络有多块，共识就有多块。

## 术语

- 压测端：压力测试中发送压力的客户端程序
- Batch：一批交易，包含 ID 和交易列表
- 交易池：接受压测端的交易，打包为 Batch，为共识服务

## 流程

Go 版本交易池启动后会监听 GRPC 端口。

压测端也是用 Go 写的，通过 AddTxs 接口发送交易列表，长度为 10000。

交易池通过 AddTxs 接收到交易列表后，生成 Batch，保存 Batch。通过 SendBatch 接口将 Batch 发送给其他交易池。

当足够多的 Batch 被其他交易池收到后，此交易池通过 AckBatch 通知其他交易池 Batch 准备好了，可以交给共识了。

交易池收到 AckBatch 后，将 Batch 保存到本地。

目前 Batch ID 是通过时间戳生成的，后续可以改为哈希。Batch 的储存规则是 `/tmp/extpool/进程PID/BatchID` 。比如交易池的进程 PID 是 1234，Batch ID 是 1613190000，那么 Batch 会存储在 `/tmp/extpool/1234/1613190000` 文件中。

区块链节点初始化时，会启动交易池进程，并且通过环境变量的方式将配置文件中的交易池监听端口、其他交易池地址 传给交易池。

主节点的共识模块向交易池拿交易的时候，会读取 `/tmp/extpool/进程PID` 下文件列表。如果文件列表为空，就认为交易池是空的；如果文件列表不为空，就取一个文件名作为 Batch ID，进行后续共识。交易池还要维护 Batch ID 的 Set，防止重复拿取 Batch。

省流版：对交易池来说，要做的就是收到交易，打包成 Batch 并传播。确保足够多的交易池有这个 Batch 后，向特定文件路径写入 Batch，让共识模块去拿。

### TODO

C++ 版本交易池的流程和 Go 版本一样，但是一些正确性相关的功能还没有实现。

- 实现 ack 机制，即 Batch 被 Quorum 个交易池拥有后，共识模块才能从交易池拿到这个 Batch。目前 Go 版本通过 AckBatch 实现了这个功能，而 C++ 版本收到 Batch 就写入文件系统。如果不实现 ack 机制，就存在正确性问题，即共识拿到了 Batch ID，但因为部分交易池宕机，导致后续节点的交易池始终无法取得此 Batch；实现 ack 机制后会增加延迟，但不怎么影响吞吐。

- 缺失交易拉取。目前在 C++ 版本中，如果在某个 Batch 广播途中，有一个交易池宕机了，那么这个交易池就永远无法拿到这个 Batch。在 Go 中，当此共识节点需要这个 Batch 的时候，会写入 `/tmp/extpool/1234/fetching/1613190000`。交易池一直监听 `/tmp/extpool/1234/fetching` 目录下的文件列表。如果有文件，就说明有共识节点需要这个 Batch，交易池就会主动向其他交易池请求这个 Batch。

- 实现单笔交易的接收，并且打包多笔交易为 Batch。从而让这个交易池可以在所有场景（功能测试、性能测试）中使用。

此外，还可以考虑：

- 现在借助文件系统实现交易池和共识之间的通信。以后可以改为共享内存等更高性能的进程间通信方式。

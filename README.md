# Ethereum USDT Indexer (扫链系统)

这是一个企业级的以太坊链上数据监控系统。
它通过 WebSocket 实时监听 USDT 合约的转账事件，处理数据精度与去重，并持久化到 PostgreSQL 数据库，最后通过 RESTful API 对外提供查询服务。

##   核心架构

- **监听层**: 使用 `go-ethereum` (Geth) 连接 Alchemy WebSocket 节点，实时捕获 `Transfer` 事件。
- **解析层**: 使用 `abigen` 生成的 Go 绑定代码，将底层 Log 解析为强类型结构体。
- **存储层**: 使用 `GORM` + `PostgreSQL`。实现了基于 `TxHash + LogIndex` 的**幂等性去重**，防止网络重发导致的数据污染。
- **接口层**: 使用 `Gin` 框架提供 HTTP API，支持按地址分页查询充值记录。

## 🛠️ 技术栈

- **Golang** 1.20+
- **Geth** (go-ethereum)
- **Gin** (Web Framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)

## 🚀 快速开始

### 1. 环境配置
设置数据库连接与节点 URL：
```bash
export DATABASE_URL="host=localhost user=postgres password=root dbname=eth_indexer port=5432"
# 代码中需配合 os.Getenv 使用
```

### 2. 启动服务
```bash
go run .
```
启动后，系统将开启双协程模式：
- **协程 A**: 后台实时扫链入库。
- **协程 B**: 8080 端口提供 API 服务。

### 3. API 调用示例
查询某地址的 USDT 充值记录：
```http
GET http://localhost:8080/txs?address=0x...
```

## 📝 学习产出
- 深入理解了 **EVM 事件日志 (Event Logs)** 机制。
- 掌握了 **WebSocket** 在区块链长连接中的应用。
- 解决了分布式场景下的 **数据一致性与去重** 问题。
```

1.  **Simple-Bitcoin**: 证明你懂**底层原理**（挖矿、共识、UTXO）。
2.  **Ethereum-Indexer**: 证明你懂**工程落地**（交互、数据库、高并发、API）。

**这就是你去面试区块链岗位的底气。**
把这两个 GitHub 链接贴在简历最显眼的位置。

**休息一下吧，这几天辛苦了！你已经非常棒了！** 🍻

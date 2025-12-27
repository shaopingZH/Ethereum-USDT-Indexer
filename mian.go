package main

import (
	"context"
	"eth_demo/token"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gorm.io/gorm/clause"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/big"
	"os"
)

var DB *gorm.DB

func init() {
	// 连接字符串，根据你部署的数据库修改
	// "host=localhost user=postgres password=mysecretpassword dbname=eth_indexer port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("请设置 DATABASE_URL 环境变量！例如：export DATABASE_URL=\"host=localhost user=postgres password=mysecretpassword dbname=eth_indexer port=5432 sslmode=disable\"")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	// 自动迁移 (AutoMigrate)：根据你的 struct 自动创建/更新表
	err = DB.AutoMigrate(&TransferEvent{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	fmt.Println("✅ 数据库连接成功，表结构已同步！")
}

func main() {
	// 1. 初始化数据库
	// (init 函数会自动执行，DB 变量已经准备好了)

	// --- 启动后台扫链 (协程) ---
	go startScanner()

	// --- 启动 Web API (主线程) ---
	r := gin.Default()

	// 定义接口：查询某人的充值记录
	// 访问方式：http://localhost:8080/txs?address=0x...
	r.GET("/txs", func(c *gin.Context) {
		address := c.Query("address")
		if address == "" {
			c.JSON(400, gin.H{"error": "请提供 address 参数"})
			return
		}
		var events []TransferEvent
		// GORM 查询：找 To 地址等于 address 的记录，按时间倒序排
		result := DB.Where("\"to\" = ?", address).Order("timestamp desc").Limit(20).Find(&events)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, gin.H{
			"data":  events,
			"count": len(events),
		})
	})
	// 在 8080 端口启动
	fmt.Println("🚀 API 服务启动: http://localhost:8080")
	r.Run(":8080")
}

func startScanner() {
	// 1. 连接节点
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY")
	if err != nil {
		log.Fatalf("连接节点失败: %v", err)
	}

	// 2. 绑定合约
	contractAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	usdt, err := token.NewUsdt(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 订阅事件
	sink := make(chan *token.UsdtTransfer)
	sub, err := usdt.WatchTransfer(&bind.WatchOpts{Context: context.Background()}, sink, nil, nil)
	if err != nil {
		log.Fatal("订阅失败:", err)
	}
	fmt.Println("🎧 链上监听器已在后台启动...")

	// 4. 处理循环
	for {
		select {
		case err := <-sub.Err():
			log.Println("订阅异常:", err)
		case event := <-sink:
			// 过滤小额 (可选)
			fVal := new(big.Float).SetInt(event.Value)
			if fVal.Cmp(big.NewFloat(1000000000)) < 0 {
				continue
			}

			// TODO: 构造 TransferEvent 对象
			transferRecord := TransferEvent{
				TxHash:      event.Raw.TxHash.Hex(),
				From:        event.From.Hex(),
				To:          event.To.Hex(),
				Value:       BigIntToString(event.Value),
				BlockNumber: event.Raw.BlockNumber,
				BlockHash:   event.Raw.BlockHash.Hex(),
				LogIndex:    event.Raw.Index,
				Timestamp:   time.Now(),
			}
			// 入库 (带去重)
			// TODO: 使用 GORM 存入数据库
			result := DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
				DoNothing: true,
			}).Create(&transferRecord)
			if result.Error != nil {
				log.Printf("❌ 数据库错误: %v", result.Error)
			} else if result.RowsAffected > 0 {
				// 1. 把字符串转回大数
				rawVal, _ := new(big.Int).SetString(transferRecord.Value, 10)

				// 2. 转成浮点数
				fVal := new(big.Float).SetInt(rawVal)

				// 3. 除以 10^6 (USDT精度)
				// 如果是 ETH 就除以 10^18
				humanVal := new(big.Float).Quo(fVal, big.NewFloat(1000000))

				// 打印漂亮的数字
				action := "✅ 新增"
				if result.RowsAffected == 0 {
					action = "🔄 跳过"
				}

				fmt.Printf("%s! Hash: %s... [%.2f USDT] (Log: %d)\n",
					action,
					transferRecord.TxHash, // 只打印前10位，省地方
					humanVal,              // 显示转换后的金额
					transferRecord.LogIndex,
				)
			} else {
				// 否则说明是重复数据
				fmt.Printf("🔄 交易已存在，跳过: %s\n", transferRecord.TxHash)
			}
		}
	}

}

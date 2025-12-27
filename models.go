package main

import (
	"math/big"
	"time"

	"gorm.io/gorm" // 👈 引入 GORM
)

// TransferEvent 是数据库中的一条转账记录
// 这里的 `gorm:"..."` 就是数据库表的字段标签，非常重要！
type TransferEvent struct {
	gorm.Model            // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	TxHash      string    `gorm:"index:idx_tx_log,unique"` // 交易哈希，唯一索引，不允许为空
	From        string    `gorm:"index"`                   // 转账方地址，建索引方便查询
	To          string    `gorm:"index"`                   // 接收方地址，建索引方便查询
	Value       string    `gorm:"type:numeric(38,0)"`      // 金额 (用字符串存大数，确保精度不丢)
	BlockNumber uint64    `gorm:"index"`                   // 区块号
	BlockHash   string    // 区块哈希
	LogIndex    uint      `gorm:"index:idx_tx_log,unique"` // 日志索引，区分同一交易中的不同事件
	Timestamp   time.Time // 交易发生时间
}

// 辅助方法：将 big.Int 转换为 string，避免精度丢失
func BigIntToString(b *big.Int) string {
	if b == nil {
		return "0"
	}
	return b.String()
}

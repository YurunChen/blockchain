package main

// 1.定义交易结构
type Transaction struct {
	TXID     []byte
	TXInput  []TXInput
	TXOutput []TXOutput
}
type TXInput struct {
	//	1.引用的交易ID
	TXid []byte
	//	2.引用的output索引值
	Index int64
	//  3.解锁脚本
	Sig string
}
type TXOutput struct {
	//转账金额
	value float64
	//锁定脚本，用地址模拟
	PubKeyHash string
}

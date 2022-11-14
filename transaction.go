package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

const reward = 12.5

// Transactions 1.定义交易结构
type Transactions struct {
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
	Value float64
	//锁定脚本，用地址模拟
	PubKeyHash string
}

func (tx *Transactions) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		return
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}
func NewCoinBaseTX(address string, data string) *Transactions {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	//对于挖矿交易来说，只有一个input和output
	transaction := Transactions{[]byte{}, []TXInput{input}, []TXOutput{output}}
	transaction.SetHash()
	return &transaction
}

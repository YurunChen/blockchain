package main

import "crypto/sha256"

type Block struct {
	//	1.前区块哈希
	PrevBlockHash []byte
	//	2.当前区块哈希
	CurBlockHash []byte
	//	3.数据
	Data []byte
}

func NewBlock(PreBlockHash []byte, Data string) *Block {
	block := &Block{
		PrevBlockHash: PreBlockHash,
		CurBlockHash:  []byte{}, //先填空，之后再计算
		Data:          []byte(Data),
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {

	blockinfo := append(block.PrevBlockHash, block.Data...)
	hash := sha256.Sum256(blockinfo)
	block.CurBlockHash = hash[:]
}

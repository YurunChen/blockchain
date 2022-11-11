package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	//	1.前区块哈希
	PrevBlockHash []byte
	//	2.当前区块哈希
	CurBlockHash []byte
	//	3.数据
	Data []byte
}

type BlockChain struct {
	blocks []*Block
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
func NewBlockChain() *BlockChain {
	gensisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{gensisBlock},
	}
}
func GenesisBlock() *Block {
	return NewBlock([]byte{}, "一期创世区块")

}
func (block *Block) SetHash() {

	blockinfo := append(block.PrevBlockHash, block.Data...)
	hash := sha256.Sum256(blockinfo)
	block.CurBlockHash = hash[:]
}
func main() {
	blockChain := NewBlockChain()
	for i, block := range blockChain.blocks {
		fmt.Printf("===============当前区块高度%d==============\n", i)
		fmt.Printf("前区块哈希值:%x\n", block.PrevBlockHash)
		fmt.Printf("当前区块哈希值:%x\n", block.CurBlockHash)
		fmt.Printf("数据:%s\n", block.Data)
	}
}

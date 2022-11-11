package main

import (
	"crypto/sha256"
	"time"
)

type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevBlockHash []byte
	//3.默克尔根
	MerkleRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机值
	Nonce uint64
	//7.当前区块哈希
	CurBlockHash []byte
	//8.数据
	Data []byte
}

func NewBlock(PreBlockHash []byte, Data string) *Block {
	block := &Block{
		Version:       00,
		PrevBlockHash: PreBlockHash,
		MerkleRoot:    []byte{},
		TimeStamp:     uint64(time.Now().Unix()),
		CurBlockHash:  []byte{}, //先填空，之后再计算
		Data:          []byte(Data),
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	var blockinfo []byte
	blockinfo = append(blockinfo, Uint64ToByte(block.Version)...)
	blockinfo = append(blockinfo, block.PrevBlockHash...)
	blockinfo = append(blockinfo, block.MerkleRoot...)
	blockinfo = append(blockinfo, Uint64ToByte(block.TimeStamp)...)
	blockinfo = append(blockinfo, Uint64ToByte(block.Difficulty)...)
	blockinfo = append(blockinfo, Uint64ToByte(block.Nonce)...)
	blockinfo = append(blockinfo, block.CurBlockHash...)
	blockinfo = append(blockinfo, block.Data...)
	hash := sha256.Sum256(blockinfo)
	block.CurBlockHash = hash[:]
}

// 实现一个辅助函数，实现将uint64转换成[]byte
func Uint64ToByte(data uint64) []byte {

}

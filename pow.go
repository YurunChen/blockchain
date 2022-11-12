package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := &ProofOfWork{
		block: block,
	}
	target_string := "0000100000000000000000000000000000000000000000000000000000000000"
	tmpInt := &big.Int{}
	tmpInt.SetString(target_string, 16)
	pow.target = tmpInt
	return pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//TODO
	var blockinfo []byte
	var nonce uint64
	var hash [32]byte
	//1.拼接数据
	block := pow.block
	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevBlockHash,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.CurBlockHash,
			block.Data,
		}
		//将二维的切片数组连接起来，返回一个一维的切片数组
		blockinfo = bytes.Join(tmp, []byte(""))
		//2.做哈希运算
		hash = sha256.Sum256(blockinfo)
		//3.比较哈希值，返回结果
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		result := tmpInt.Cmp(pow.target)
		//a.找到了，退出返回
		if result == -1 {
			fmt.Printf("挖矿成功！hash:%x, nonce:%d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}
}

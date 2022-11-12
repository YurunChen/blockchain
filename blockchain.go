package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB
	tail []byte
}

const BlockChainDB = "BlockChain.db"
const BlockChainBucket = "bucket"

var lastHash []byte

// 建立新区块链
func NewBlockChain() *BlockChain {
	db, err := bolt.Open(BlockChainDB, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败！")
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockChainBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BlockChainBucket))
			if err != nil {
				log.Panic("创建一个bucket!")
			}
			gensisBlock := GenesisBlock()
			bucket.Put(gensisBlock.CurBlockHash, gensisBlock.toByte())
			bucket.Put([]byte("lastHashKey"), gensisBlock.CurBlockHash)
		} else {
			lastHash = bucket.Get([]byte("lastKeyHash"))
		}
		return nil
	})
	return &BlockChain{
		db,
		lastHash,
	}
}
func (blockChain *BlockChain) AddBlock(data string) {
	//lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	//prevHash := lastBlock.CurBlockHash
	//block := NewBlock(prevHash, data)
	//blockChain.blocks = append(blockChain.blocks, block)
}

// 创世区块
func GenesisBlock() *Block {
	return NewBlock([]byte{}, "一期创世区块")
}

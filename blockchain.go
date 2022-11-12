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

// 建立新区块链
func NewBlockChain() *BlockChain {
	db, err := bolt.Open(BlockChainDB, 0600, nil)
	//defer db.Close()
	var lastHash []byte
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
			bucket.Put(gensisBlock.CurBlockHash, gensisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"), gensisBlock.CurBlockHash)
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})
	return &BlockChain{
		db,
		lastHash,
	}
}

func (bc *BlockChain) AddBlock(data string) {
	//lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	//prevHash := lastBlock.CurBlockHash
	//block := NewBlock(prevHash, data)
	//blockChain.blocks = append(blockChain.blocks, block)
	//如何获取前区块的哈希呢？？
	db := bc.db         //区块链数据库
	lastHash := bc.tail //最后一个区块的哈希
	db.Update(func(tx *bolt.Tx) error {
		//完成数据添加
		bucket := tx.Bucket([]byte(BlockChainBucket))
		if bucket == nil {
			log.Panic("bucket不应该为空，请检查!")
		}
		//a. 创建新的区块
		block := NewBlock(lastHash, data)
		//b. 添加到区块链db中
		//hash作为key， block的字节流作为value，尚未实现
		bucket.Put(block.CurBlockHash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.CurBlockHash)
		//c. 更新一下内存中的区块链，指的是把最后的小尾巴tail更新一下
		bc.tail = block.CurBlockHash
		return nil
	})
}

// 创世区块
func GenesisBlock() *Block {
	return NewBlock([]byte{}, "一期创世区块")
}

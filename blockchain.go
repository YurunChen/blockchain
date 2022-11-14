package main

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//4. 引入区块链
//2. BlockChain结构重写
//
//使用数据库代替数组

type BlockChain struct {
	//定一个区块链数组
	//blocks []*Block
	db *bolt.DB

	tail []byte //存储最后一个区块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

// 5. 定义一个区块链
func CreateBlockChain(address string) *BlockChain {
	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希， 从数据库中读出来的
	var lastHash []byte

	//1. 打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2. 找到抽屉bucket(如果没有，就创建）
		//没有抽屉，我们需要创建
		bucket, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			log.Panic("创建bucket(b1)失败")
		}

		//创建一个创世块，并作为第一个区块添加到区块链中
		genesisBlock := GenesisBlock(address)

		//3. 写数据
		//hash作为key， block的字节流作为value，尚未实现
		bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
		bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
		lastHash = genesisBlock.Hash

		////这是为了读数据测试，马上删掉,套路!
		//blockBytes := bucket.Get(genesisBlock.Hash)
		//block := Deserialize(blockBytes)
		//fmt.Printf("block info : %s\n", block)

		return nil
	})

	return &BlockChain{db, lastHash}
}

// 只是返回区块链实例，不创建
func NewBlockChain() *BlockChain {
	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希， 从数据库中读出来的
	var lastHash []byte

	//1. 打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要操作数据库（改写）
	db.View(func(tx *bolt.Tx) error {
		//2. 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic(err)
		}

		lastHash = bucket.Get([]byte("LastHashKey"))

		return nil
	})

	return &BlockChain{db, lastHash}
}

// 定义一个创世块
func GenesisBlock(address string) *Block {
	coinbase := NewCoinBaseTX(address, "一期创世块")
	return NewBlock([]*Transactions{coinbase}, []byte{})
}

// 5. 添加区块
func (bc *BlockChain) AddBlock(transaction []*Transactions) {
	//如何获取前区块的哈希呢？？
	db := bc.db         //区块链数据库
	lastHash := bc.tail //最后一个区块的哈希

	db.Update(func(tx *bolt.Tx) error {

		//完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空，请检查!")
		}

		//a. 创建新的区块
		block := NewBlock(transaction, lastHash)

		//b. 添加到区块链db中
		//hash作为key， block的字节流作为value，尚未实现
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		//c. 更新一下内存中的区块链，指的是把最后的小尾巴tail更新一下
		bc.tail = block.Hash

		return nil
	})
}

func (bc *BlockChain) Printchain() {

	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key-> value 进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)
			//fmt.Printf("key=%x, value=%s\n", k, v)
			fmt.Printf("=============== 区块高度: %d ==============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transaction[0].TXInput[0].Sig)
			return nil
		})
		return nil
	})
}
func (bc *BlockChain) FindUTXO(address string) []TXOutput {
	var UTXO []TXOutput
	//TODO
	return UTXO
}

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
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
		Difficulty:    0,
		Nonce:         0,
		CurBlockHash:  []byte{}, //先填空，之后再计算
		Data:          []byte(Data),
	}
	//block.SetHash()
	proofOfWork := NewProofOfWork(block)
	hash, nonce := proofOfWork.Run()
	//查找目标的随机数，不断的哈希运算
	block.CurBlockHash = hash
	//根据挖矿解决对区块数据进行更新
	block.Nonce = nonce
	return block
}

//
//func (block *Block) SetHash() {
//	//var blockinfo []byte
//	//blockinfo = append(blockinfo, Uint64ToByte(block.Version)...)
//	//blockinfo = append(blockinfo, block.PrevBlockHash...)
//	//blockinfo = append(blockinfo, block.MerkleRoot...)
//	//blockinfo = append(blockinfo, Uint64ToByte(block.TimeStamp)...)
//	//blockinfo = append(blockinfo, Uint64ToByte(block.Difficulty)...)
//	//blockinfo = append(blockinfo, Uint64ToByte(block.Nonce)...)
//	//blockinfo = append(blockinfo, block.CurBlockHash...)
//	//blockinfo = append(blockinfo, block.Data...)
//	//tmp := [][]byte{
//	//	Uint64ToByte(block.Version),
//	//	block.PrevBlockHash,
//	//	Uint64ToByte(block.TimeStamp),
//	//	Uint64ToByte(block.Difficulty),
//	//	Uint64ToByte(block.Nonce),
//	//	block.CurBlockHash,
//	//	block.Data,
//	//}
//	////将二维的切片数组连接起来，返回一个一维的切片数组
//	//blockinfo = bytes.Join(tmp, []byte(""))
//	//hash := sha256.Sum256(blockinfo)
//	//block.CurBlockHash = hash[:]
//}

// 实现一个辅助函数，实现将uint64转换成[]byte
func Uint64ToByte(data uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
func (block *Block) toByte() []byte {
	//TODO
	return []byte{}
}

// 序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	//- 使用gob进行序列化（编码）得到字节流
	//1. 定义一个编码器
	//2. 使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错!")
	}

	//fmt.Printf("编码后的小明：%v\n", buffer.Bytes())

	return buffer.Bytes()
}

// 反序列化
func Deserialize(data []byte) Block {

	decoder := gob.NewDecoder(bytes.NewReader(data))

	var block Block
	//2. 使用解码器进行解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错!", err)
	}

	return block
}

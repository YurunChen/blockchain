package main

import "fmt"

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("1111111111")
	blockChain.AddBlock("2222222222")
	it := blockChain.NewIterator()
	for {
		block := it.Next()
		fmt.Println("================================")
		fmt.Printf("前区块哈希：%x\n", block.PrevBlockHash)
		fmt.Printf("当前区块哈希：%x\n", block.CurBlockHash)
		fmt.Printf("当前区块哈希：%s\n", block.Data)
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	//for i, block := range blockChain.blocks {
	//	fmt.Printf("===============当前区块高度%d==============\n", i)
	//	fmt.Printf("前区块哈希值:%x\n", block.PrevBlockHash)
	//	fmt.Printf("当前区块哈希值:%x\n", block.CurBlockHash)
	//	fmt.Printf("数据:%s\n", block.Data)
	//}
}

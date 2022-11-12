package main

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("区块2")
	blockChain.AddBlock("区块3")
	//for i, block := range blockChain.blocks {
	//	fmt.Printf("===============当前区块高度%d==============\n", i)
	//	fmt.Printf("前区块哈希值:%x\n", block.PrevBlockHash)
	//	fmt.Printf("当前区块哈希值:%x\n", block.CurBlockHash)
	//	fmt.Printf("数据:%s\n", block.Data)
	//}
}

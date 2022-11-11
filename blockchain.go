package main

type BlockChain struct {
	blocks []*Block
}

// 建立新区块链
func NewBlockChain() *BlockChain {
	gensisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{gensisBlock},
	}
}
func (blockChain *BlockChain) AddBlock(data string) {
	lastBlock := blockChain.blocks[len(blockChain.blocks)-1]
	prevHash := lastBlock.CurBlockHash
	block := NewBlock(prevHash, data)
	blockChain.blocks = append(blockChain.blocks, block)
}

// 创世区块
func GenesisBlock() *Block {
	return NewBlock([]byte{}, "一期创世区块")
}

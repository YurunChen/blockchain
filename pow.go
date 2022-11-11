package main

import (
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
	target_string := "0000010000000000000000000000000000000"
	tmpInt := &big.Int{}
	tmpInt.SetString(target_string, 16)
	pow.target = tmpInt
	return pow
}

func (ProofOfWork *ProofOfWork) Run() ([]byte, uint64) {
	//TODO
	return []byte("hello world"), 10
}

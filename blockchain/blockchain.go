package blockchain

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	Blocks []*Block
}

func (c *BlockChain) AddBlock(data string) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (c *BlockChain) ToString() string {
	out := "Block Info :"
	var elems []string
	for _, b := range c.Blocks {
		elems = append(elems, fmt.Sprintf("\n{\n\tPrevious Hash:\t%x\n\tData:\t%s\n\tHash:\t%x\n}", b.PrevHash, b.Data, b.Hash))
	}

	return out + strings.Join(elems, ",")
}

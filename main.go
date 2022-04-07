package main

import (
	"fmt"
	"github.com/ashans/go-chain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	fmt.Println(chain.ToString())
}

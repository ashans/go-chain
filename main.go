package main

import (
	"github.com/ashans/go-chain/blockchain"
	"github.com/ashans/go-chain/cli"
	"github.com/ashans/go-chain/errors"
	"github.com/dgraph-io/badger"
	"os"
)

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer func(Database *badger.DB) {
		err := Database.Close()
		errors.Handle(err)
	}(chain.Database)

	command := cli.NewCommandLine(chain)
	command.Run()
}

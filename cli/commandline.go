package cli

import (
	"flag"
	"fmt"
	"github.com/ashans/go-chain/blockchain"
	"github.com/ashans/go-chain/errors"
	"os"
	"runtime"
	"strings"
)

type CommandLine struct {
	chain *blockchain.BlockChain
}

func NewCommandLine(c *blockchain.BlockChain) *CommandLine {
	return &CommandLine{c}
}

func (c *CommandLine) printUsage() {
	fmt.Println("Usage :")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (c *CommandLine) addBlock(data string) {
	c.chain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (c *CommandLine) printChain() {
	iter := c.chain.Iterator()

	out := "Blockchain Info :"
	var elems []string

	for {
		b := iter.Next()
		elems = append(elems, fmt.Sprintf("\n{\n\tPrevious Hash:\t%x\n\tData:\t%s\n\tHash:\t%x\n}", b.PrevHash, b.Data, b.Hash))

		if len(b.PrevHash) == 0 {
			break
		}
	}

	fmt.Print(out + strings.Join(elems, ","))
}

func (c *CommandLine) Run() {
	c.validateArgs()

	addBlockCommand := flag.NewFlagSet("add", flag.ExitOnError)
	printCommand := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCommand.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCommand.Parse(os.Args[2:])
		errors.Handle(err)
	case "print":
		err := printCommand.Parse(os.Args[2:])
		errors.Handle(err)
	default:
		c.printUsage()
		runtime.Goexit()
	}
	if addBlockCommand.Parsed() {
		if *addBlockData == "" {
			addBlockCommand.Usage()
			runtime.Goexit()
		}
		c.addBlock(*addBlockData)
	}
	if printCommand.Parsed() {
		c.printChain()
	}
}

func (c *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		runtime.Goexit()
	}
}

package main

import (
    "fmt"
    "github.com/aditya999123/BlockChain-101/blockchain"
    "strconv"
    "flag"
    "os"
    "runtime"
)

type CommandLine struct {
    blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
    fmt.Println("Usage:")
    fmt.Println(" add -block BLOCK_DATA - add a block to the chain:")
    fmt.Println(" print - Prints the block in the chain")

}

func (cli *CommandLine) validateArgs() {
    if len(os.Args) < 2 {
        cli.printUsage()
        runtime.Goexit()
    }
}

func (cli *CommandLine) printChain() {
    iter := cli.blockchain.Iterator()

    for {
        block := iter.Next()
        fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Data in block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := blockchain.NewProof(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        if len(block.PrevHash) == 0 {
            break
        }
    }
}

func (cli *CommandLine) run() {
    cli.validateArgs()

    addBlockCmd   := flag.NewFlagSet("add", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
    addBlockData := addBlockCmd.String("block", "", "Block Data")

    switch os.Args[1] {
        case "add":
            err := addBlockCmd.Parse(os.Args[2:])
            blockchain.Handle(err)

        case "print":
            err := printChainCmd.Parse(os.Args[2:])
            blockchain.Handle(err)

        default:
            cli.printUsage()
            runtime.Goexit()
    }

    if addBlockCmd.Parsed() {
        if *addBlockData == "" {
            addBlockCmd.Usage()
            runtime.Goexit()
        }

        cli.blockchain.AddBlock(*addBlockData)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }
}

func main() {
    defer os.Exit(0)
    chain := blockchain.InitBlockChain()

    defer chain.Database.Close()

    cli := CommandLine{chain}
    cli.run()
}
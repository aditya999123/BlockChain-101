package main

import (
    "fmt"
    "github.com/aditya999123/BlockChain-101/blockchain"
)

func main() {
    chain := blockchain.InitBlockChain()

    chain.AddBlock("First block after genisis")
    chain.AddBlock("Second block after genisis")
    chain.AddBlock("Third block after genisis")

    for _, block := range chain.Blocks {
        // fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Data in block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
    }
}
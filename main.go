package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
)

type Block struct {
    Hash     []byte
    Data     []byte
    PrevHash []byte 
}

func (b *Block) DeriveHash() {
    info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
    block := &Block{[]byte{}, []byte(data), prevHash}
    block.DeriveHash()
    return block
}

type BlockChain struct {
    blocks []*Block
}

func (chain *BlockChain) AddBlock (data string) {
    prevBlock := chain.blocks[len(chain.blocks)-1]
    newBlock := CreateBlock(data, prevBlock.Hash)
    chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
    return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
    return &BlockChain{[]*Block{Genesis()}}
}

func main() {
    chain := InitBlockChain()

    chain.AddBlock("First block after genisis")
    chain.AddBlock("Second block after genisis")
    chain.AddBlock("Third block after genisis")

    for _, block := range chain.blocks {
        // fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Data in block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
    }
}
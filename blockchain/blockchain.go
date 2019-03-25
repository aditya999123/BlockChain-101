package blockchain

import (
	"github.com/dgraph-io/badger"
	"fmt"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
    LastHash []byte
    Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			
			genesis := Genesis()
			fmt.Println("Genesis proved")
			
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
		
			lastHash = genesis.Hash

			err = txn.Set([]byte("lh"), genesis.Hash)
			return err
		} else {
			lastHash, err = item.Value()
			return err
		}
 	})

 	Handle(err)

 	blockchain := BlockChain{lastHash, db}
 	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})

	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)

		err = txn.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash

		return err
	})

	Handle(err)
}


func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var encodedData []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
	
		encodedData, err = item.Value()

		return err
	})

	Handle(err)
	block := Deserialize(encodedData)
	iter.CurrentHash = block.PrevHash

	return block
}
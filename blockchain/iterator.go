package blockchain

import (
	"github.com/dgraph-io/badger"
	"log"
)

type BlockIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (i *BlockIterator) Next() *Block {

	var block *Block
	err := i.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(i.CurrentHash)
		if err != nil {
			log.Panic(err)
		}
		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})

		return err
	})
	if err != nil {
		log.Panic(err)
	}

	i.CurrentHash = block.PrevHash

	return block
}

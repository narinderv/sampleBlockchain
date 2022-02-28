package main

import (
	"crypto/sha256"
	"fmt"

	//	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         string
	hash         string
	previousHash string
	timestamp    time.Time
	PoW          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (blk Block) generateHash() string {
	blockData := blk.data + blk.previousHash + blk.timestamp.String() + strconv.Itoa(blk.PoW)

	blockHash := sha256.Sum256([]byte(blockData))

	return fmt.Sprintf("%x", blockHash)
}

func (blk *Block) mineBlock(difficulty int) {
	for !strings.HasPrefix(blk.hash, strings.Repeat("0", difficulty)) {
		blk.PoW++
		blk.hash = blk.generateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		data:      "",
		hash:      "0",
		timestamp: time.Now(),
		PoW:       0,
	}

	return Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
	}
}

func (blkChain *Blockchain) addNewBlock(info string) {
	lastBlock := blkChain.chain[len(blkChain.chain)-1]

	newBlock := Block{
		data:         info,
		hash:         "",
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
		PoW:          0,
	}

	newBlock.mineBlock(blkChain.difficulty)

	fmt.Printf("Data: %s\nHash: %s\nPrev. Hash: %s\nTimestamp: %s\nPoW: %d\n", newBlock.data, newBlock.hash, newBlock.previousHash, newBlock.timestamp.String(), newBlock.PoW)

	blkChain.chain = append(blkChain.chain, newBlock)
}

func (blkChain Blockchain) validateBlockchain() bool {
	for i := range blkChain.chain[1:] {
		curBlock := blkChain.chain[i+1]
		prevBlock := blkChain.chain[i]

		if curBlock.hash != curBlock.generateHash() || curBlock.previousHash != prevBlock.hash {
			return false
		}
	}

	return true
}

func main() {
	blockChain := CreateBlockchain(3)

	blockChain.addNewBlock("Block1")
	blockChain.addNewBlock("Block2")

	fmt.Println(blockChain.validateBlockchain())
}

package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	nonce        int
	transaction  string
	previousHash string
	currentHash  string
}

type BlockChain struct {
	list []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := new(Block)
	block.transaction = transaction
	block.nonce = nonce
	block.previousHash = previousHash
	block.currentHash = CalculateHash(transaction + strconv.Itoa(nonce) + previousHash)
	return block
}

func (bk *BlockChain) AddBlock(block *Block) {
	bk.list = append(bk.list, block)
}

func ListBlocks(bk *BlockChain) {
	for index, block := range bk.list {
		fmt.Println("Block: ", index+1)
		fmt.Println("Transaction: ", block.transaction)
		fmt.Println("Nonce: ", block.nonce)
		fmt.Println("Current Hash: ", block.currentHash)
		fmt.Println("Previous Hash: ", block.previousHash)
		fmt.Println()
	}
}

func (bk *BlockChain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bk.list) {
		bk.list[index].transaction = newTransaction
		bk.list[index].currentHash = CalculateHash(newTransaction + strconv.Itoa(bk.list[index].nonce) + bk.list[index].previousHash)
	}
}

func (bk *BlockChain) VerifyChain() bool {
	for i := 1; i < len(bk.list); i++ {

		currentBlock := bk.list[i]
		previousBlock := bk.list[i-1]

		currentHash := CalculateHash(currentBlock.transaction + strconv.Itoa(currentBlock.nonce) + currentBlock.previousHash)
		if currentHash != currentBlock.currentHash {
			return false
		}

		if previousBlock.currentHash != currentBlock.previousHash {
			return false
		}
	}

	return true
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

func NewBlockChain() *BlockChain {
	blockchain := &BlockChain{}
	return blockchain
}

func (bk *BlockChain) GetHeadBlockCurrentHash() string {
	if len(bk.list) > 0 {
		return bk.list[len(bk.list)-1].currentHash
	}
	return ""
}

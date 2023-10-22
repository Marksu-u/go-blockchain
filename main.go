package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

var Blockchain []Block

// calculateHash generates a SHA-256 hash for a block based on its data
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generateBlock creates a new block
func generateBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().Format(time.RFC3339)
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

// isBlockValid checks if the block hasn't been tampered with
func isBlockValid(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func main() {
	// Create the "genesis block"
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "Genesis Block - The beginning of the chain",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	
	// Add the genesis block to our blockchain
	Blockchain = append(Blockchain, genesisBlock)
	fmt.Println("Genesis block created!")

	// Let's add some data (e.g., simulating transactions)
	transactions := []string{
		"Alice sends 5 coins to Bob",
		"Bob sends 2 coins to Charlie",
		"Charlie sends 1 coin to Alice",
	}

	for _, tx := range transactions {
		lastBlock := Blockchain[len(Blockchain)-1]
		
		newBlock, _ := generateBlock(lastBlock, tx)
		
		if isBlockValid(newBlock, lastBlock) {
			Blockchain = append(Blockchain, newBlock)
			fmt.Printf("Block #%d created successfully!\n", newBlock.Index)
		}
	}

	// Print out entire blockchain in readable JSON format
	bytes, _ := json.MarshalIndent(Blockchain, "", "  ")
	fmt.Println("\nOur final Blockchain:")
	fmt.Printf("%s\n", bytes)
}

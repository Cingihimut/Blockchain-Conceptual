package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func genesisBlock() *Block {
	return generateBlock(nil, "genesis Block")
}

func generateBlock(prevBlock *Block, data string) *Block {
	block := &Block{
		Timestamp: time.Now(),
		Data:      data,
	}
	if prevBlock != nil {
		block.PrevHash = prevBlock.Hash
	}

	//PoW (dummy implementation)
	for nonce := 0; ; nonce++ {
		block.Nonce = nonce
		block.Hash = calculateHash(block)
		if isValidHash(block.Hash) {
			break
		}
	}
	return block
}

func calculateHash(block *Block) string {
	record := string(block.Timestamp.Unix()) + block.Data + block.PrevHash + string(block.Nonce)
	hash := sha256.New()
	hash.Write([]byte(record))
	hased := hash.Sum(nil)
	return hex.EncodeToString(hased)
}

func isValidHash(hash string) bool {
	return hash[:3] == "000"
}

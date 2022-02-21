package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(data + newBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	newBlock.hash = hexHash
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Prev Hash: %s\n", block.prevHash)
		fmt.Printf("Hash: %s\n", block.hash)
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis block")
	chain.addBlock("Second block")
	chain.addBlock("Third block")
	chain.listBlocks()
}

// blockchain struct를 구현하여 block의 array를 만든다.
// reciever 함수를 사용하여, 해당 struct에서 사용할 수 있는 메소드를 생성한다.

// addBlock은 data를 받아 block을 추가하는 메소드이다.
// 새 block을 형성하여 hash와 prevHash를 추가한다.
// 이전 강의에서 달라진 것은 없다

// listBlocks는 모든 block을 보여준다
// Printf는 format을 맞추어서 print해주는 함수이다.

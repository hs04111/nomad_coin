package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.PrevHash + b.Data))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

// 리팩토링
// 1. sync 패키지를 이용하여 blockchain 형성을 단 한번만 한다.
// gorutine이 많이 시행되거나, 병렬적으로 프로그램이 처리되어도
// blockchain은 하나만 형성될 것이다.

// 2. blockchain에는 block의 pointer를 넣을 것이다.
// blockchain이 길어질수록 데이터의 양이 방대해지기 때문

// 3. calculateHash 함수는 reciever 함수로, 해당 객체의 값을
// mutate할 수 있다.

// 4. AddBlock과 AllBlock은 밖에서도 사용할 수 있도록 첫글자 대문자로 생성

// 4-1 참고로 AllBlock은 reciever 함수가 아니어도 된다.
// GetBlockchain().block을 리턴하면 된다. main.go에서 사용법이 달라진다.

// 5. main.go에서 block의 정보에 접근할 수 있도록 type에서 첫글자를 대문자로 변경.

// 전반적으로 좀 더 module을 생성하는 방식으로 작업중

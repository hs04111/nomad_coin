package main

import (
	"fmt"

	"github.com/hs04111/nomad_coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}

// 리팩토링: Singleton pattern을 이용한다
// Singleton pattern이란? 한 어플리케이션에서 하나의 인스턴스만 사용하는 패턴이다
// 이를 위해서, blockchain package에서 객체의 생성을 컨트롤한다.
// main.go에서는 blockchain 인스턴스를 형성할 방법은 GetBlockchain 함수를 사용하는 방법 뿐이다.

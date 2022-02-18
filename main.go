package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data string
	hash string
	prevHash string
}

func main(){
	genesisBlock := block{"Genesis Block", "",""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash	
}

// block은 data와 hash를 가지며, hash는 이전 block의 hash와
// data를 더해 해쉬함수를 적용한 값이다. prevHash가 이전 block
// 의 hash이다.

// genesisBlock은 첫 번째 block이다. prevHash는 빈 값이 되어야 한다.
// hash 함수로  SHA256 알고리즘이 대부분 사용된다. go는 내장함수로 가지고 있다.
// sha256.Sum256은 byte의 slice를 인자로 받는다. string을 []byte
// 로 바꾸는 방법은 위와 같다.

// 다만 sha256의 리턴값도 []byte인데, 보통의 암호화폐의 경우 이를
// 16진법으로 hash값을 쓴다. fmt.Springf는 첫 번째 인자의 형식으로
// 해당 값을 변환하여 리턴한다. 이를 genesis block의 hash에 넣는다.
package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}

// GetBlockchain에서 인스턴스를 형성한다.
// 다만, 변수를 선언한 후 nil인 경우에만 인스턴스를 형성한다.
// 그 외에는 기존의 객체를 사용해야 한다.
// 오직 이 package에서만 객체를 형성하는 것을 컨트롤할 수 있다.

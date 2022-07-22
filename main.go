package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hs04111/nomad_coin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle  string
	Blockchain []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(rw, homeData{"Home", blockchain.GetBlockchain().AllBlocks()})

}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// go에서는 기본 패키지에서 서버를 열 수 있다.
// http.ListenAndServe의 첫번째 인자는 포트 주소이고, 두번째는 핸들러
// log.Fatal은 내부에서 에러가 있으면 Fatal이 작동한다.

// HandleFunc는 Node js의 route와 같다.
// handler function은 인자로 rw, r이 있는데 rw는 응답을 작성하는 곳
// r은 요청(request)의 포인터이다. 요청이 클 수도 있으므로, 포인터만 가져오는 것

// rw는 writer로, 여기에 무언가를 쓰려면 Fprint를 사용하여 writer에
// 작성하여야 한다.

// template를 렌더링하려면 html/template package를 사용해야 한다.
// ParseFiles로 html을 가져오고, Excute를 통하여 데이터를 template로 줄 수 있다.
// data는 struct로 정의하고, field는 대문자로 쓰여야 template에서 사용 가능하다.
// 함수들의 인자와 리턴하는 값들을 확인할 것. error를 반환하기도 한다.
// teplate.Must는 err를 자동으로 처리해준다. 컨트롤 클릭하여 코드 확인해볼 것.
// blockchain은 반드시 GetBlockchain으로만 가져와야 한다. 그렇게 할 수 밖에 없다.

package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello!")
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
// r은 요청의 포인터이다. 요청이 클 수도 있으므로, 포인터만 가져오는 것

// rw는 writer로, 여기에 무언가를 쓰려면 Fprint를 사용하여 writer에
// 작성하여야 한다.

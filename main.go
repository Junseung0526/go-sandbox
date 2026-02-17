package main

import "fmt"

func main() {
	// calculator.go에 있는 함수를 별도 임포트 없이 바로 쓸 수 있음

	//기존 방식
	var name string = "gri22ly"

	//go 스타일
	age := 24
	name2 := "junseung"

	fmt.Println(name, age, name2)

	//선언과 동시에 초기화 (:= 사용)
	scores := map[string]int{
		"Kim":  90,
		"Lee":  85,
		"Park": 100,
	}

	//값 추가 및 수정
	scores["Choi"] = 95

	//값 삭제
	delete(scores, "Lee")

	//데이터 확인
	//존재 여부를 같이 알려줌
	val, ok := scores["Kim"]
	if ok {
		fmt.Println("찾은 값:", val)
	}
}

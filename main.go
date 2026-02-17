package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Dept     string
	IsActive bool
}

func (s Student) Introduce() {
	fmt.Printf("안녕하세요, %s학과 %d살 %s입니다!\n", s.Dept, s.Age, s.Name)
}

//포인터로 데이터를 메모리 주소로 넘겨줌
func (s *Student) HaveBirthday() {
	s.Age++
}

func main() {
	// --- Map 연습 파트 ---
	prt := map[string]int{
		"junseung": 24,
		"hyojun":   21,
		"minseung": 22,
	}

	val, ok := prt["grizzly"]
	if ok {
		fmt.Println("나이:", val)
	} else {
		fmt.Println("데이터가 존재하지 않습니다.")
	}

	// --- 구조체 연습 파트 ---
	s1 := Student{
		Name: "김준승",
		Age:  24,
		Dept: "스마트소프트웨어",
	}

	s2 := Student{"홍길동", 20, "컴퓨터공학", true}

	// 출력 테스트
	fmt.Println("s1 정보:", s1.Name, s1.Age)
	fmt.Println("s2 정보:", s2.Name, s2.Age)

	s1.HaveBirthday()
	s2.HaveBirthday()

	fmt.Println("s1 정보:", s1.Name, s1.Age)
	fmt.Println("s2 정보:", s2.Name, s2.Age)

	// 메서드 호출 테스트
	s1.Introduce()
}

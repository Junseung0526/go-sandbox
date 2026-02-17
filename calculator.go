package main

import "fmt"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("0으로 나눌 수 없습니다")
    }
    return a / b, nil
}
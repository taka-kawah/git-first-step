package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := 2
	switch p {
	case 1:
		p1()
	case 2:
		p2()
	default:
		fmt.Println("おわり！")
	}
}

func p1() {
	f := 1.11
	numf := int(f)
	fmt.Println(numf)
	fmt.Println("numfの型は...", reflect.TypeOf(numf))

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}

func p2() {
	//Q1. 以下のスライスから最も小さい数を返すコード
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	temp := l[0]
	for i := 1; i < len(l); i++ {
		if temp > l[i] {
			temp = l[i]
		}
	}
	fmt.Printf("Q1の答えは%v", temp)

	fmt.Println()

	//Q2. 以下の果物の価格の合計を出力するコードを書いてください。
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	fmt.Printf("Q2の答えは%v", sum)
}

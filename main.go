package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {
	p := 7
	switch p {
	case 1:
		p1()
	case 2:
		p2()
	case 3:
		p3()
	case 4:
		p4()
	case 5:
		p5()
	case 6:
		p6()
	case 7:
		p7()
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

type Vertex struct {
	x, y int
}

func (v Vertex) Plus() int {
	return v.x + v.y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.x, v.y)
}

func p3() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}

func goroutine(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func p4() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

func producer(first chan<- int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		fmt.Println("firstチャンネルに入れる", i)
		first <- i
	}
}

func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		temp := i * 2
		fmt.Println("secondチャンネルに入れる", temp)
		second <- temp
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		temp := i * 4
		fmt.Println("thirdチャンネルに入れる", temp)
		third <- temp
	}
}

func p5() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)

	for r := range third {
		fmt.Println(r)
	}
}

func gr1(c chan string) {
	for {
		//ネットワークからくるパケットを延々と取得し続けるイメージ
		c <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func gr2(c chan string) {
	for {
		//ネットワークからくるパケットを延々と取得し続けるイメージ
		c <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func p6() {
	defer fmt.Println("END")
	c1 := make(chan string)
	c2 := make(chan string)
	go gr1(c1)
	go gr2(c2)

	timeout := time.After(10 * time.Second)

	for {
		//同時に到着したパケットを分けて取得
		select {
		case msg1 := <-(c1):
			fmt.Println(msg1)
		case msg2 := <-(c2):
			fmt.Println(msg2)
		case <-timeout:
			fmt.Println("time has come...")
			return
		default:
			fmt.Println("Quiet...")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	//専有
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) getv(key string) int {
	//専有
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func p7() {
	c := Counter{v: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(&c, c.getv("key"))
}

package main

import (
	"fmt"
	"./zoo/animals"
	"runtime"
	"time"
)
var S=""

func init(){
	S=S + "A"
}
func init(){
	S=S + "B"
}
func init(){
	S=S + "C"
}

func main()  {

	fmt.Println("Hello World!")

	//fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

	go fmt.Println("Yeah")
	fmt.Printf("NumCPU:%d¥r¥n", runtime.NumCPU())
	fmt.Printf("NumGoroutine:%d¥n", runtime.NumGoroutine())
	fmt.Printf("Version:%s¥n", runtime.Version())
	fmt.Println()

	// init関数の動作確認⇒ABC
	fmt.Println(S)
	// スライスの確認
	s:=make([]int, 10)
	fmt.Println(s)

	//var a [10]int
	//fmt.Println(a)

	a:=make([]float64, 3)
	fmt.Println(a)
	a[0]=3.14
	fmt.Println(a)
	a[1]=6.28
	fmt.Println(a)
	lenval:=make([]int,8)
	lenlen := len(lenval)
	fmt.Println(lenlen)

	lena := [3]int{1,2,3}
	lenalen:=len(lena)
	fmt.Println(lenalen)

ch :=make(chan int, 20)
go recieve("1st goroutine", ch)
go recieve("2st goroutine", ch)
go recieve("3st goroutine", ch)

i := 0
	for i < 100 {
		ch<- i
		i++
	}
	close(ch)

// ゴルーチンの完了を三秒待つ
time.Sleep(3 * time.Second)
}

func recieve (name string, ch <-chan int){
	for  {
		i, ok := <-ch
		if ok ==false{
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)

	}
	fmt.Println(name + "is done.")
}
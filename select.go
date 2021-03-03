package main

import (
	"fmt"
	"time"
)

func fibo(n int) (res int) {

	if n <= 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func printfibo(res int) {
	fmt.Println("fibonacci : ", res)
}

func concur(f func(int), temp chan int) {

	for {

		select {
		case n := <-temp:
			fmt.Printf("\b")
			f(n)
			return
		default:
			for _, c := range `-\|/` {
				fmt.Printf("\b%c", c)
				time.Sleep(80 * time.Millisecond)
			}
		}

	}
}

func main() {
	num := make(chan int)

	go func() {
		num <- fibo(40)
	}()

	concur(printfibo, num)

	return

}

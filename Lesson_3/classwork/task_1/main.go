package main

import "fmt"

var (
	N int = 10
	K int = 3
)

func main() {
	for i := 0; i < N; i++ {
		if i%K == 0 {
			fmt.Println(i)
		}
	}
}

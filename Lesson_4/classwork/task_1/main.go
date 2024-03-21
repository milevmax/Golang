package main

import "fmt"

func print(sl []int) {
	fmt.Println("Len = ", len(sl), "Cap = ", cap(sl))
}

func main() {
	var power int = 2
	s := make([]int, 10)
	for i, _ := range s {
		s[i] = 2 * power
		power = s[i]
		fmt.Println(s)
		print(s)
	}
	for i := 0; i < 40; i++ {
		print(s)
		s = append(s, i)
	}
	fmt.Println(s)
}

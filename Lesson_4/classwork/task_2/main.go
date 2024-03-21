package main

import "fmt"

func avg(s []int) float32 {
	count_marks := len(s)

	var sum_mark int
	for _, mark := range s {
		sum_mark += mark
	}
	var avr_mark float32 = float32(sum_mark) / float32(count_marks)
	return avr_mark
}

func main() {
	marks := []int{9, 4, 4, 7, 10, 10}
	//count_marks := len(marks)
	//
	//var sum_mark int
	//for _, mark := range marks {
	//	sum_mark += mark
	//}
	//var avr_mark float32 = float32(sum_mark) / float32(count_marks)
	//fmt.Println("average mark is ", avr_mark)

	avg_marks := avg(marks)
	fmt.Println("average mark is ", avg_marks)
}

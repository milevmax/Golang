package main

import (
	"fmt"
	"reflect"
)

func main() {
	var first int = 10
	fmt.Println("first number is", first)

	var second int
	fmt.Println("second number is", second)

	var (
		third int = 10
	)

	second = third * 2

	fmt.Println("second number is", second, "third number is", third)
	fmt.Println(reflect.TypeOf(third))

	const (
		my_year int = 1998
		month       = 6
	)

	type status bool
	type yearOfBirthd = int16

	var (
		classic_status   bool         = true
		custom_status    status       = false
		classic_birthday int16        = 1998
		custom_birthday  yearOfBirthd = 1998
	)

	fmt.Println("buildin status value:", classic_status, "buildin status type:", reflect.TypeOf(classic_status), "\n",
		"custom status value:", custom_status, "custom status type:", reflect.TypeOf(custom_status), "\n",
		"buildin birthday value:", classic_birthday, "buildin birthday type:", reflect.TypeOf(classic_birthday), "\n",
		"custom birthday value:", custom_birthday, "custom birthday type:", reflect.TypeOf(custom_birthday))
}

//func biography(age)

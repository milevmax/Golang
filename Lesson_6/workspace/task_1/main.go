package main

import "fmt"

type smallInterface interface {
	foo()
}

type smallStruct struct {
	smallNumber int
}

func (s smallStruct) foo() {
	fmt.Println("foo")
}

type bigInterface interface {
	smallInterface
	click()
}

type bigStruct struct {
	bigNumber int
}

func (b bigStruct) foo() {
	fmt.Println("foo")
}

func (b bigStruct) click() {
	fmt.Println("click click click")
}

type myInterfaces []smallInterface

func (m *myInterfaces) addToMyInterface(s smallInterface) {
	*m = append(*m, s)
}

type EmbedNumber1 struct {
	ownNumber int
	bigStruct
}

type EmbedNumber2 struct {
	ownNumber  int
	bigStruct2 bigStruct
}

func main() {
	smallStructInstance := smallStruct{}
	bigStructInstance := bigStruct{}
	mis := make(myInterfaces, 0)
	mis.addToMyInterface(smallStructInstance)
	fmt.Println(mis)
	mis.addToMyInterface(bigStructInstance)
	fmt.Println(mis)

	en1 := EmbedNumber1{
		ownNumber: 10,
		bigStruct: bigStruct{bigNumber: 20},
	}

	fmt.Println(en1)

	en2 := EmbedNumber2{
		ownNumber:  10,
		bigStruct2: bigStruct{bigNumber: 123},
	}

	fmt.Println(en2)
}

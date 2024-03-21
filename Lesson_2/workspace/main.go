package main

import "fmt"

type person struct {
	num  int
	name string
}

func (p person) describe() string {
	return fmt.Sprintf("person name=%v", p.name)
}

type house struct {
	address string
	person
}

type car struct {
	model  string
	master *person
}

func (c *car) sellCar() {
	c.master = nil
}

func main() {
	me := person{num: 1, name: "Max"}
	var p_me *person
	p_me = &me

	fmt.Println(me)

	//my_house := house{
	//	address: "aboba 228",
	//	person: person{num: 1,
	//		name: "Shrek"}}
	//fmt.Println(me.describe())
	//fmt.Println(my_house.describe())
	//
	//bb := 1 > 0
	//fmt.Println(bb)

	myCar := car{
		model:  "subaru",
		master: p_me}

	fmt.Println(myCar)

	myCar.sellCar()

	fmt.Println(myCar)
}

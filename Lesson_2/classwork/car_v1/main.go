package main

import "fmt"

type Car struct {
	model string
	color string
}

type Parking struct {
	car      Car
	occupied bool
}

func (c Car) ShowInfo() {

	fmt.Printf("Car model is %v and car color is %v\n", c.model, c.color)
}

func (p Parking) ShowInfo() {
	fmt.Printf("parking place occupied: %v", p.occupied)
}

func main() {
	car := Car{color: "black", model: "Honda"}
	parking := Parking{car: car, occupied: true}

	car.ShowInfo()
	parking.ShowInfo()
}

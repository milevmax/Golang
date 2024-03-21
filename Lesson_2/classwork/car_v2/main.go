package main

import (
	"fmt"
	"time"
)

type Car struct {
	color       string
	modelNumber int
}

type Parking struct {
	Spot              *Car
	timeStatusChanged string
}

func (c Car) showCarInfo() string {
	return fmt.Sprintf("Car color is %v, model Num: %d", c.color, c.modelNumber)
}

func placeCar(c *Car, p *Parking) {
	if p.Spot == nil {
		p.Spot = c
		fmt.Println("place occupied by ", c)
	} else {
		fmt.Println("place already occupied by ", p.Spot)
	}
	//fmt.Println(p.Spot)
}

//func freePlace !!!!

func main() {
	car := Car{color: "black", modelNumber: 1}
	fmt.Printf(car.showCarInfo())

	dt := time.Now()
	var dt_formeted string = dt.Format("2006-01-02 15:04:05")
	fmt.Println()
	fmt.Println(dt_formeted)

	parking := Parking{}
	fmt.Println()
	fmt.Println("kek", parking.timeStatusChanged, parking.Spot)

	p_car := &car

	//parking = Parking{Spot: p_car}
	fmt.Printf("The type of myVar is %T\n", p_car)
	fmt.Printf("The type of myVar is %T\n", car)
	p_parking := &parking
	placeCar(p_car, p_parking)
	placeCar(p_car, p_parking)
}

package main

import "fmt"

type Transport struct {
	driver string
}

func (t *Transport) downTheControls(driver string) {
	t.driver = driver
}

func (t *Transport) showDriver() {
	fmt.Println("Driver is ", t.driver)
}

type Car struct {
	Transport
	countWheels int
}

func (c *Car) hurryUp() {
	if c.countWheels == 4 {
		fmt.Println("Врум Врум!")
	} else {
		fmt.Println("Пук сренкь =(")
	}
}

func main() {
	t1 := Transport{}
	t1.showDriver()

	var bestDriver string = "Ryan Gosling"
	t1.downTheControls(bestDriver)
	t1.showDriver()
	fmt.Println()

	var worstDriver string = "Сергій Самовілов"
	car := Car{countWheels: 4}
	car.downTheControls(worstDriver)
	car.showDriver()
	car.hurryUp()
	fmt.Println()

	var catDriver string = "Веталь"
	bike := Car{countWheels: 2}
	bike.downTheControls(catDriver)
	bike.showDriver()
	bike.hurryUp()
	fmt.Println()
}

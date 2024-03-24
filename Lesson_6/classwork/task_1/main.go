package main

import (
	"classworkTask1/forecast"
	"fmt"
)

type tomorrowTemperature interface {
	GetTomorrowTemperature() int
}

func printForecast(tt tomorrowTemperature) {
	temperatureForecast := tt.GetTomorrowTemperature()
	fmt.Println(temperatureForecast)
}

func main() {
	//defaultForecast := forecast.ConstantForecast{DefaultTemperatureValue: 7}
	//t := defaultForecast.GetTomorrowTemperature()
	//fmt.Println(t)
	printForecast(forecast.ConstantForecast{10})
	printForecast(forecast.RandomForecast{})
	printForecast(forecast.WetherAPIForecast{"https://wttr.in/Kyiv?format=3"})
}

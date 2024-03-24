package forecast

import (
	"io"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
)

type ConstantForecast struct {
	DefaultTemperatureValue int
}

func (cf ConstantForecast) GetTomorrowTemperature() int {
	return cf.DefaultTemperatureValue
}

type RandomForecast struct {
}

func (rf RandomForecast) GetTomorrowTemperature() int {
	return rand.IntN(50) - 25
}

type WetherAPIForecast struct {
	APIurl string
}

func (af WetherAPIForecast) GetTomorrowTemperature() int {
	resp, _ := http.Get(af.APIurl)
	body, _ := io.ReadAll(resp.Body)
	response := string(body)

	// Regular expression to match the temperature part
	re := regexp.MustCompile(`\+?(-?\d+)°C`)
	matches := re.FindStringSubmatch(response)
	temperature, _ := strconv.Atoi(matches[1])
	//fmt.Println(string(body))
	return temperature
}

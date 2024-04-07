package transport

import "fmt"

func printHoursAndMinutes(hours float32) {
	wholeHours := int(hours)                      // Get the whole number of hours
	minutes := (hours - float32(wholeHours)) * 60 // Convert the remaining fraction of an hour to minutes

	// Print the result with formatting
	fmt.Printf("Being on the way Hours: %d, Minutes: %d\n", wholeHours, int(minutes))
}

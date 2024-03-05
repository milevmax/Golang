package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	fullName      string = "Jackie Chan"
	chineseName   string = "成龍"
	yearOfBirth   int    = 1954
	mounthOfBirth string = "April"
)

var (
	currentYear int = time.Now().Year()
	age             = uint8(currentYear - yearOfBirth)
	firstName       = strings.Split(fullName, " ")[0]
)

type profession string

func (p profession) String() string {
	return fmt.Sprintf("profession is %s", string(p))
}

func main() {
	var person_profession profession = "actor"
	fmt.Println(person_profession, age, firstName)
	//fmt.Printf("%v model is %v and car color is %v\n", c.model, c.color)
	fmt.Printf("%T\n\n", "pharmacist")
	fmt.Printf("%s was born in %d of %s, so now he is %d years old. His Chinese name is %s.\n%ss %s. ",
		fullName, yearOfBirth, mounthOfBirth, age, chineseName, firstName, person_profession)
	fmt.Print("during his career he performed a large number of dangerous stunts and always maintained a good physical shape.")
}

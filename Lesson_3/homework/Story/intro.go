package Story

import "fmt"

func InitIntro() Intro {
	introGame := Intro{IntroText: "Ти прокинувся в вагоні потягу. Світло не горить, схоже на вагон метро.\nНе пам'ятаєш як тут опинився. Твої очі звикають до темряви..."}
	return introGame
}

type Intro struct {
	IntroText   string
	GameStarted bool
}

func (i *Intro) TakeInput() {
	fmt.Println("Введіть R щоб почати:")

	var input string

	fmt.Scanf("%s", &input)
	if input == "R" {
		i.GameStarted = true
	} else {
		fmt.Print("GG")
	}
}

func (i Intro) HandleInput() (updatedState string) {
	if i.GameStarted {
		fmt.Println(i.IntroText)
	}
	updatedState = CarriageStateName
	return updatedState
}

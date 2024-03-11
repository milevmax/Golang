package main

import (
	"fmt"
	"math/rand"
)

//
//Основна ідея:
// 1) сутність клітка, поля: довжина, ширина float, наявність води, зайнятість bool.
//	Методи: розрахувать об'єм, отримати статус зайнятості.
// 2) сутність тварина, поля: назва, категорія, вага
//3) сутність наглядач, поля: ім'я, досвід.  методи: перевести звіра в клітку(ім'я звіра, номер клітки)
//4) функція пристасовоності до клітки (клітка, звір)

const squareWeightFraction float32 = 1 / 5

type WaterEnv struct {
	waterEnv bool
}

//type Holder struct {
//	holdAnimal *Animal
//}

type Cage struct {
	number     int
	width      float32
	height     float32
	occupiedBy *Animal
	WaterEnv
}

func (c Cage) calcSquare() float32 {
	return c.height * c.width
}

type Animal struct {
	breed string
	//name  string
	weigh float32
	cage  *Cage
	WaterEnv
}

type ZooKeeper struct {
	name       string
	experience int
	hold       *Animal
}

func (z *ZooKeeper) catchAnimal(animal *Animal) {
	if z.experience <= 5 && 0.5 > rand.Float32() {
		fmt.Printf("%v couldn't catch the %v right now, try again!\n", z.name, animal.breed)
	} else {
		z.hold = animal
		//fmt.Println(z)
		if animal.cage != nil {
			animal.cage.occupiedBy = nil
			animal.cage = nil
		}
		fmt.Printf("%v caught the %v!\n", z.name, animal.breed)
	}
}

func (z *ZooKeeper) putToCage(cage *Cage) {

	switch {
	case z.hold == nil:
		fmt.Println("You didn't catch anyone\n")
	case cage.occupiedBy != nil:
		fmt.Printf("The cage occupied by %v\n", cage.occupiedBy)
	case cage.calcSquare()/z.hold.weigh < squareWeightFraction:
		fmt.Printf("The cage is too small for %v\n", z.hold.breed)
	case z.hold.waterEnv != cage.waterEnv:
		fmt.Printf("cage %v doesn`t fit for %v\n", cage.number, z.hold.breed)
	default:
		z.hold.cage = cage
		cage.occupiedBy = z.hold
		//z.hold = nil
		fmt.Printf("%v got into the cage number: %v\n", z.hold.breed, cage.number)
	}
}

func main() {
	tiger_1 := Animal{breed: "tiger", weigh: 400, WaterEnv: WaterEnv{waterEnv: false}}

	//tiger_2 := Animal{breed: "tiger", weigh: 50, WaterEnv: WaterEnv{waterEnv: false}}
	//aligator := Animal{breed: "aligator", weigh: 200, WaterEnv: WaterEnv{waterEnv: true}}
	//frog := Animal{breed: "frog", weigh: 0.4, WaterEnv: WaterEnv{waterEnv: true}}
	//bobre := Animal{breed: "bobre kurva", weigh: 6, WaterEnv: WaterEnv{waterEnv: true}}
	//
	cage_1 := Cage{number: 1, width: 10, height: 10, WaterEnv: WaterEnv{waterEnv: true}}
	//cage_2 := Cage{number: 2, width: 20, height: 12, WaterEnv: WaterEnv{waterEnv: true}}
	//cage_3 := Cage{number: 3, width: 5, height: 2, WaterEnv: WaterEnv{waterEnv: true}}
	cage_4 := Cage{number: 4, width: 25, height: 25, WaterEnv: WaterEnv{waterEnv: false}}
	//cage_5 := Cage{number: 5, width: 10, height: 50, WaterEnv: WaterEnv{waterEnv: false}}
	//
	keeper_1 := ZooKeeper{name: "Petro", experience: 22}
	//keeper_2 := ZooKeeper{name: "Bob", experience: 3}

	keeper_1.catchAnimal(&tiger_1)
	keeper_1.putToCage(&cage_1)

	keeper_1.putToCage(&cage_4)
}

package main

import (
	"fmt"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (a *Animal) New(eat, move, speak string) {
	a.eat = eat
	a.move = move
	a.speak = speak
}

func (a *Animal) Eat() {
	fmt.Println(a.eat)
}

func (a *Animal) Move() {
	fmt.Println(a.move)
}

func (a *Animal) Speak() {
	fmt.Println(a.speak)
}

func main() {

	var cow Animal
	var snake Animal
	var bird Animal

	cow.New("grass", "walk", "moo")
	bird.New("worms", "fly", "peep")
	snake.New("mice", "slither", "hsss")

	animals := []Animal{
		cow,
		snake,
		bird,
	}

	for {
		var animal, info string
		fmt.Print(">")
		fmt.Scanf("%s %s", &animal, &info)
		var animalSelected Animal
		switch string(animal) {
		case "cow":
			animalSelected = animals[0]
		case "bird":
			animalSelected = animals[1]
		case "snake":
			animalSelected = animals[2]
		}
		switch string(info) {
		case "eat":
			animalSelected.Eat()
		case "move":
			animalSelected.Move()
		case "speak":
			animalSelected.Speak()
		}
	}
}
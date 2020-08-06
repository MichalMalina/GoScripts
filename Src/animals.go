package main

import "fmt"
import "strings"
import "bufio"
import "os"

type Animal interface {
        Eat()
        Move()
        Speak()
}

type Cow struct {
        name string
}

type Snake struct {
        name string
}

type Bird struct {
        name string
}

func (c *Cow) Eat() {
        fmt.Println("grass")
}

func (c *Cow) Move() {
        fmt.Println("walk")
}

func (c *Cow) Speak() {
        fmt.Println("moo")
}

func (b *Bird) Eat() {
        fmt.Println("worms")
}

func (b *Bird) Move() {
        fmt.Println("fly")
}

func (b *Bird) Speak() {
        fmt.Println("peep")
}

func (s *Snake) Eat() {
        fmt.Println("mice")
}

func (s *Snake) Move() {
        fmt.Println("slither")
}

func (s *Snake) Speak() {
        fmt.Println("hsss")
}

func main() {

        var animals = make(map[string]Animal)

        for {
                fmt.Printf(">")
                scanner := bufio.NewScanner(os.Stdin)
                scanner.Scan()
                strs := strings.Split(scanner.Text(), " ")
                if len(strs) < 3 {
                        fmt.Println("Wrong number of elements in input.")
                        continue
                }
                cm1 := strs[0]
                cm2 := strs[1]
                cm3 := strs[2]

                if cm1 == "newanimal" {

                        switch cm3 {
                        case "cow":
                                animals[cm2] = new(Cow)
                        case "snake":
                                animals[cm2] = new(Snake)
                        case "bird":
                                animals[cm2] = new(Bird)
                        default:
                                fmt.Printf("Wrong animal type")
                        }
                } else if cm1 == "query" {
                        _, present := animals[cm2]
                        if present {
                                switch cm3 {
                                case "move":
                                        (animals[cm2]).Move()
                                case "eat":
                                        (animals[cm2]).Eat()
                                case "speak":
                                        (animals[cm2]).Speak()
                                default:
                                        fmt.Printf("Wrong action.")
                                }
                        } else {
                                fmt.Printf("Animal name doesnt exist")
                        }

                } else {
                        fmt.Printf("Wrong command please start with query or newanimal ")
                }
        }
}


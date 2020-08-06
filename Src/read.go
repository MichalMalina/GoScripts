package main

import "fmt"
import "strings"
import "io/ioutil"

func main() {
	var inputer string
	type Person struct {
						fname string
						lname string
						}
						var list []Person

						fmt.Printf("Put in name of file to open: ")
						fmt.Scan(&inputer)
						content, err := ioutil.ReadFile(inputer)
						if err != nil {
							fmt.Println(err)
						}
						for _, line := range strings.Split(strings.TrimSuffix(string(content), "\n"), "\n") {
							var per Person
							for i,name := range strings.Fields(line) {
							if i == 0 {
							per.fname = name
							} else {
							per.lname = name
							}
							}
							list = append(list, per)
						}

						for _, person := range list{
								fmt.Println(person.fname + " " + person.lname)
							}
}
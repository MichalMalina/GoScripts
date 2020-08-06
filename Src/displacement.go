package main

import "fmt"

func main() {
//Global vars and scans
var acel float64
fmt.Printf("Please enter the aceleration: ")
fmt.Scan(&acel)

var vel float64
fmt.Printf("Please enter the velocity: ")
fmt.Scan(&vel)

var init float64
fmt.Printf("Please enter the initial diplacement: ")
fmt.Scan(&init)

var time float64
fmt.Printf("Please enter time: ")
fmt.Scan(&time)

//Run the functions
first:= GenDisplaceFn(acel, vel , init)
fmt.Printf("The result is: ")
fmt.Println(first(time))

}

//Calcutating functions
func GenDisplaceFn( acel, vel, init float64) func(time float64) float64 {
fn:= func (time float64) float64 {
return 0.5*acel*time*time + vel*time + init
}
return fn
}
package main

import "fmt"

func main() {

var something float32
var somethingElse float32

	fmt.Printf("Welcome\n")

	var x int = 1
	
	switch x {
		
		case 1:
		
		fmt.Printf("Please put in a number: ")

		fmt.Scanln(&something)

		fmt.Println("truncated to: " , int(something))
		
		x++
	
			switch x {
			
				case 2:

				fmt.Printf("Please put in another decimal: ")

				fmt.Scanln(&somethingElse)

				fmt.Println("truncated to: " , int(somethingElse))
				
				x++
			
					switch x {
						
						case 3:
						
						fmt.Println("Thanks thats enougth your integers are: " , int(something) , "and" , int(somethingElse))
					
					}
			}
	}
}


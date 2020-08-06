package main

import "fmt"

func main() {
        var x int = 1
        var y int = 2

		// Random addition to x and y variables
        calc1(x, y)
		// If statement prints different strings depending on x and y values
        calc2(x, y)
}

func calc1(x, y int) {
        var inputer int
        fmt.Printf("Put in a random int: ")
        _, err := fmt.Scan(&inputer)
        if err != nil {
                fmt.Printf("Inputed value is not an integer \n")
        }

        x = x + inputer + y
        fmt.Printf("Your number is:")
        fmt.Println(x + y)

}

func calc2(x, y int) {
        if x == 1 && y == 1 {
                fmt.Printf("NO")
        } else {
                fmt.Printf("YES")
        }

}



//Race condition is when the outcome of some tasks is dependent on the time and order the tasks were executed. 
//In this case the value of the x and y variable depends on the order of execution therefore if the code is not run from a to z different strings will be printed
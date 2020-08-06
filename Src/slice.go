package main

import (
	"fmt"
	"strings"
	"sort"
	"strconv"
	
)

func main(){

	
	sli := make([]int, 3)
	
	var inputer string
	
	fmt.Printf("Enter some number: ")
	fmt.Scan(&inputer)
	inputer = strings.ToLower(inputer)

	for inputer != "x"{
	x, err := strconv.Atoi(inputer)
		if err == nil {
			sli = append(sli, x)
	}
	sort.Ints(sli)
	fmt.Println(sli[3:])
	
		fmt.Scan(&inputer)
		inputer = strings.ToLower(inputer)
		}
		

	sort.Ints(sli)
	fmt.Println(sli[3:])

}

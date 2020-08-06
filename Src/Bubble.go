package main
import "fmt"

 
func main() {
	// Global data
	var inputer int 
	theSlice := []int {}
	
	//Initial Prompt
	fmt.Printf("Please put in a number and hit enter ten times \n")
    fmt.Scan(&inputer)
    
	//Loop Scan
	for i := 0 ; i < 10 ; i++ {
	 fmt.Scan(&inputer)
	 theSlice = append(theSlice , inputer)
   }
   
   //Run bubble sort
   BubbleSort(theSlice)
   
   //Print sorted slice
   fmt.Println(theSlice)
}

	//sort numbers
	func BubbleSort(numbers []int) {

	   for i := len(numbers); i > 0; i-- {

		Swap(numbers , i)
		  }
		 
	 }
	 
	 //swap numbers
	 func Swap( numbers []int, i int ) {
		 for j := 1; j < i; j++ {
			 if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
							}
				 }
	}

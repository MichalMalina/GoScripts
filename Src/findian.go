package main 


import "fmt"
import "strings"
import "strconv"

func main () {


  var x int
  var y *int
  z := 3
  y = &z
  x = &y


  var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)


  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)



  i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)



var inputer string

fmt.Printf("Pleaese put in a string: ")
fmt.Scanln(&inputer)

if strings.HasPrefix(inputer , "i") && strings.Contains(inputer , "a") && strings.HasSuffix(inputer , "n") {fmt.Printf("Found!")} else {fmt.Printf("Not found.")}

}
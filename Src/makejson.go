package main

import "fmt"
import "encoding/json"

func main() {
var name string
var adress string

fmt.Printf("Your name: ")
fmt.Scan(&name)
fmt.Printf("Your address: ")
fmt.Scan(&adress)

commits := map[string]string {
    name:adress,
}

for key, value := range commits {
    fmt.Println("Key:", key, "Value:", value)
}

barr ,err := json.Marshal(commits)
if err == nil {
fmt.Println(string(barr))
}


}
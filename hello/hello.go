package main

import "fmt"

//Hello : get hellow string
func Hello(name string) string {
	return "Hello, " + name
}
func main(){
	fmt.Println(Hello("Chris"))
}
package main

import "fmt"

func init()  {
	fmt.Println("Init!")
}
func bazz(){
	fmt.Println("Buzz")
}

func main()  {
	bazz()
	fmt.Println("Hello world!")
}
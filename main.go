package main

import (
	"fmt"
	basic "golang-tutorial/basicOperations"
)

func main() {
	fmt.Println("This is main!")
	var list = []float64{3, 4, 5, 1, 10, 9}
	basic.SayHello()
	fmt.Println("maximum number : ", basic.Findmax(list))
	fmt.Println("average : ", basic.Average(list))

}

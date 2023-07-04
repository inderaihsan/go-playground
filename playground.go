package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func binarysearch(list []int, value int) int {
	var left = list[0]
	var right = list[len(list)-1]
	for left < right {
		var mid = (left + right) / 2
		if list[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if list[left] == value {
		return 1
	} else {
		return 0
	}
}
func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("every star brought you to tears again")
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var x = numbers[2]
	var y = numbers[2]
	fmt.Println(binarysearch(numbers, 100))
	fmt.Println(numbers)
	if add(x, y) <= 10 {
		fmt.Println("numbers are less than or equal to 10", add(x, y))
	} else {
		fmt.Println("numbers greater than 10")
	}
}

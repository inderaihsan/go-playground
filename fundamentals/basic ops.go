package main

import (
	"fmt"
	"math"
	"strings"
)

func findmax(array []float64) float64 {
	result_max := float64(0)
	for i := 0; i < len(array); i++ {
		if array[i] > float64(result_max) {
			result_max = float64(array[i])
		}
	}
	return result_max
}
func printTriangle(degree_a, degree_b, degree_c float64) {
	degree_list := []float64{degree_a, degree_b, degree_c}
	maxDegree := findmax(degree_list)
	for i := 1; i <= int(maxDegree); i++ {
		line := ""
		if i <= int(degree_a) {
			line += strings.Repeat("*", i) + " "
		}
		if i <= int(degree_b) {
			line += strings.Repeat("*", i) + " "
		}
		if i <= int(degree_c) {
			line += strings.Repeat("*", i)
		}
		fmt.Println(line)
	}
}

func triangle_degree(a int, b int, c int) (float64, float64, float64) {
	degree_a := math.Cos(float64((b ^ 2 + c ^ 2 - a ^ 2)) / (2 * float64(b*c)))
	degree_b := math.Cos(float64((a ^ 2 + c ^ 2 - b ^ 2)) / (2 * float64(a*c)))
	degree_c := math.Cos(float64((a ^ 2 + b ^ 2 - c ^ 2)) / (2 * float64(a*b)))
	return degree_a, degree_b, degree_c
}

func main() {
	fmt.Println("This is how you make an output in Golang")
	a := 5
	b := 4
	c := 3
	degree_a, degree_b, degree_c := triangle_degree(a, b, c)
	fmt.Println("Degrees:", degree_a, degree_b, degree_c)
	if degree_a > 90 || degree_b > 90 || degree_c > 90 {
		fmt.Println("one of the angle is having a degree more than 90")
		fmt.Println("Degree A : ", degree_a)
		fmt.Println("Degree B : ", degree_b)
		fmt.Println("Degree C : ", degree_c)
	}

	printTriangle(float64(a), float64(b), float64(c))
}

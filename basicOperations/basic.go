package basicOperations

import (
	"fmt"
)

func SayHello() {
	fmt.Println("HEY!")
}

func Findmax(array []float64) float64 {
	result_max := float64(0)
	for i := 0; i < len(array); i++ {
		if array[i] > float64(result_max) {
			result_max = float64(array[i])
		}
	}
	return result_max
}

func FindMin(array []float64) float64 {
	result_min := float64(0)
	for i := 0; i < len(array); i++ {
		if array[i] < result_min {
			result_min = array[i]
		}
	}
	return result_min
}
func Average(array []float64) float64 {
	count := len(array)
	sum := float64(0)
	for i := 0; i < count; i++ {
		sum = float64(sum) + array[i]
	}
	result := sum / float64(count)
	return float64(result)
}

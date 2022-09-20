package main

import (
	"fmt"
)

// Roll our own 'Ordering' interface
// Define a Type set MyOrdered
type MyOrdered interface {
	string | int | float64
}

// unconstrained template T any
func printGeneric[T any](slice []T) {
	for _, v := range slice {
		fmt.Printf("%v, ", v)
	}
}

// Polymorphic func: 'largest'
func largest[T MyOrdered](slice []T) T {
	largest := slice[0]
	for _, val := range slice {
		if largest < val {
			largest = val
		}
	}
	return largest
}

type MyInt int

func (a MyInt) String() string {
	return fmt.Sprintf("%d", a)
}

func main() {
	slice1 := []int{20, 3, 5, 27, 11, 13}
	fmt.Printf("Largest number  Int : %v -> ", largest(slice1))
	printGeneric(slice1)
	fmt.Print("\n")

	slice2 := []float64{-1000.20, 3.0, 5.0, 27.0, 100.01, 0.013}
	fmt.Printf("Largest number float64: %v -> ", largest(slice1))
	printGeneric(slice2)
	fmt.Print("\n")

	slice3 := []MyInt{MyInt(20), MyInt(3), MyInt(5), MyInt(27), MyInt(11), MyInt(13)}
	fmt.Printf("Largest number MyInt: %v -> ", largest(slice3))
	printGeneric(slice3)
	fmt.Print("\n")
}

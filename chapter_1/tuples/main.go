package main 

import (
	"fmt"
)

// A tuple is a finite sorted list of elements. It is a data structure that groups data. 
// Tuples are typically immutable sequential collections. 
// The element has related fields of different datatypes. 
// The only way to modify a tuple is to change the fields. 
// Operators such as + and * can be applied to tuples. 
// A database record is referred to as a tuple. 


//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (square int, cube int) {
	return a*a, a*a*a
}


func main() {
	square, cube := powerSeries(2)

	fmt.Println("Square ", square, "Cube", cube)
}
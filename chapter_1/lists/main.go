package main

import (
	"container/list"
	"fmt"
)


// A list is a sequence of elements. 
// Each element can be connected to another with a link in a forward or backward direction. 
// The element can have other payload properties. 
// This data structure is a basic type of container. 
// Lists have a variable length and developer can remove or add elements more easily than an array. 
// Data items within a list need not be contiguous in memory or on disk. Linked lists were proposed by Allen Newell, Cliff Shaw, and Herbert A. 
// Simon at RAND Corporation. 



func main(){
	var intList list.List

	intList.PushBack(3)
	intList.PushBack(4)
	intList.PushBack(6)

	for element := intList.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value.(int))
	}
}



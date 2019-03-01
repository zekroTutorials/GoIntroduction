// In dieser Datei befindet sich der Code, welchen ich "live"
// in der Aufnahme des Tutorials erstellt habe.

package main

import (
	"fmt"
)

func main() {

	// int a; -> undefined
	// int a = 5;

	var a string
	var b = "hey was geht"
	c := "hey was geht"

	fmt.Println(a, b, c)

	var myFloat float64
	pi := 3.1415

	fmt.Println(myFloat, pi)

	x, y, z, desc := 4, 7, 1, "myStuffCords"
	fmt.Println(x, y, z, desc)

	// int myArray[] = {1, 2, 3};
	var myArray []int
	// myArray[0] = 1 -> index put of range exception
	fmt.Println(len(myArray))

	myArray = append(myArray, 3)
	fmt.Println(myArray)

	var mySizedArray [5]int
	mySizedArray[3] = 2
	fmt.Println(mySizedArray)

	mySizedArray2 := make([]int, 6)
	fmt.Println(mySizedArray2)

	myPredefinedArray := []int{1, 2, 3, 4, 5}
	fmt.Println(myPredefinedArray)

	var myMultiDimensionaArray [3][3]int
	fmt.Println(myMultiDimensionaArray)

	// var myMap map[string]int
	myMap := make(map[string]int)
	myMap["test"] = 1
	fmt.Println(myMap)
}

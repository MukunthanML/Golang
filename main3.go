package main

import "fmt"

/*An array is a numbered sequence of elements of a specific length.
The type [n]T is an array of n values of type T.
An array's length is part of its type, so arrays cannot be resized.
*/

func main3() {

	//Creating with zero values
	// var a1 [3]int
	// fmt.Println(a1)
	// a1[0] = 999
	// fmt.Println(a1)
	//Creating with values
	// a2 := [3]string{"Go", "Python", "Java"}
	// fmt.Println(a2)
	//Creating with values , size as ...  three dots (…) in Golang is termed as Ellipsis
	// a3 := [...]float64{1.1, 2.1, 3.4, 4.4}
	//fmt.Println(a3)
	//len() method
	// fmt.Println(len(a3))

	//array is of value type not of reference type.
	// old := [...]int{100, 200, 300, 400, 500}
	// fmt.Println("Old array(Before):", old)
	// new := old
	// fmt.Println("New array(Before):", new)

	// Change the value at index 0 to 500
	// new[0] = 444
	// fmt.Println("New array(After):", new)
	// fmt.Println("Old array(After):", old)

	//comparision
	//same length and same type
	// fmt.Println([2]int{2, 1} == [2]int{2, 1})
	// fmt.Println([2]int{2, 1} == [2]int{2, 99})

	//different length and same type
	//fmt.Println([4]int{2, 1, 3, 5} == [2]int{2, 1})

	//same length and different type
	//fmt.Println([2]int{2, 1} == [2]float64{2.3, 1.3})

	//Multidimensional array
	mularray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(mularray)

}

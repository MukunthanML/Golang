package main

import "fmt"

/*An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
flexible view into the elements of an array. In practice,
 slices are much more common than arrays.

 A slice does not store any data, it just describes a section of an underlying array.

 []T is a slice with elements of type T
*/

func main4() {

	// Slices from Arrays
	// arr := [5]int{1, 2, 3, 4, 5}

	// fmt.Println(arr)

	// s := arr[1:4]
	// fmt.Println(s)
	// fmt.Println("**********************************************")
	// //len() and cap()

	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println("**********************************************")

	// //Slice is type of reference

	// s[0] = 888
	// fmt.Println(s)
	// fmt.Println(arr)
	// fmt.Println("**********************************************")

	// //slicing default values
	// fmt.Println(arr[:])
	// fmt.Println(arr[0:5])
	// fmt.Println(arr[:5])
	// fmt.Println(arr[0:])
	// fmt.Println("**********************************************")

	// //nill array size and len are zero
	// zeroslice := []int{}
	// fmt.Println(len(zeroslice))
	// fmt.Println(cap(zeroslice))
	// fmt.Println("**********************************************")

	// dynamically-sized arrays. with make()
	//d := make([]int, 5) // make( []T, len)
	// fmt.Println(len(d))
	// fmt.Println(cap(d))
	// fmt.Println(d)

	d := make([]int, 5, 10) // make( []T, len, cap)
	// fmt.Println(len(d))
	// fmt.Println(cap(d))
	// fmt.Println(d)

	//append to slice
	d = append(d, 22, 33)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println(d)

	/*
		If there's sufficient capacity in the underlying slice,
		the element is placed after the last element and the length get incremented.
		However, if there is not sufficient capacity, a new slice is created,
		all of the existing elements are copied over,
		the new element is added onto the end, and the new slice is returned.
	*/

	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

}

//https://www.golangprograms.com/go-language/slices-in-golang-programming.html
//https://gobyexample.com/slices
//https://tour.golang.org/moretypes/13

package main

import "fmt"

func main5() {

	a := 22       //var a int = 22
	pointer := &a //var pointer *int

	fmt.Println(a)        //prints 'a' value
	fmt.Println(&a)       //prints address of where 'a' value saved
	fmt.Println(pointer)  //pointer variable which holds the address
	fmt.Println(*pointer) //dereferencing address and gets the value from that address

	//pass by value(copy of value)
	fmt.Println("************pass by value(copy of value)******************")
	fmt.Println(a)
	passByValue(a)
	fmt.Println(a)

	fmt.Println("************pass by address******************")

	//pass by address

	fmt.Println(a)
	passByAddress(pointer) //same as passByAddress(&a)
	fmt.Println(a)

	fmt.Println("******************************")

	// fmt.Println(&pointer)  //Even pointer will have address

}

func passByValue(b int) {
	b = 999
}

func passByAddress(b *int) {
	*b = 555
}

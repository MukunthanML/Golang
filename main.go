package main

//A struct is a collection of fields.(both pre-defined, user defined)
import (
	"fmt"
	"reflect"
)

type citizen struct { //struct with pre-defined data types
	name string
	age  int
}
type employee struct { //struct with user-defined struct and pre-defined data types
	worker citizen
	salary int
}

func main() {

	//creating variable using struct, has default values
	var ram citizen
	fmt.Println(ram)
	ram.name = "Ramkumar"
	ram.age = 33
	fmt.Println(ram)

	//creating variable using struct, with values
	david := citizen{name: "David billa", age: 28}
	fmt.Println(david)

	//creating variable using struct, with new keyword
	rahim := new(citizen)
	rahim.name = "Rahim bhai"
	rahim.age = 39
	fmt.Println(rahim)
	fmt.Println(reflect.TypeOf(rahim))

	//creating pointer for struct and accessing without dereference
	//with dereference (both same in case of struct)
	gupta := &citizen{"Vivek Gupta", 33}
	fmt.Println(gupta)

	gupta.name = "Gupta raj"
	(*gupta).age = 29
	fmt.Println(gupta)

	//However, there is difference when we have assignment with
	//struct and point of struct
	// ram:=citizen{name: "Ramkumar", age: 28}
	fmt.Println("Before assignment - Ram")
	fmt.Println(ram)
	alienOne := ram  //passes a copies/values
	alienTwo := &ram //passes/shares the address
	fmt.Println("Before assignment - Alien one", alienOne)
	fmt.Println("Before assignment - Alien two", alienTwo)

	alienOne.name = "Humaniod ROBO"
	alienTwo.age = 999

	fmt.Println("After assignment - Alien one", alienOne)
	fmt.Println("After assignment - Alien two", alienTwo)
	fmt.Println("After changing aliens - Ram ", ram)

	//Lets create variable for nested struct
	bestEmployee := employee{worker: citizen{"Rajini", 18}, salary: 5000}
	smartEmployee := employee{citizen{"Kamal", 20}, 5000}
	fmt.Println(bestEmployee)
	fmt.Println(smartEmployee.worker)

}

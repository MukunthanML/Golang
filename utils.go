package utills

import "fmt"

var internal = "only inside module"

//External variable test
//comment should start with the name of the variable.
var External = "Available for all"

func internalFunc() {

	fmt.Println("internal Function")
}

//ExternalFunc is a public function.
//comment should start with the name of the function.
func ExternalFunc() {

	fmt.Println("External Function")
}

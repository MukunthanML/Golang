package main

import "fmt"

//shorthand syntax is not allowed at global
	//topic:="variables"
	//var topic string ="variables"
	var topic = 999

func test() {
	//only declaration
		//var name string
	// declaration and intialization
		 //var name string="Mukunthan"
	//Type inference
		//var name ="Mukunthan-test"
	//short hand syntax
			//name:="Mukunthan-walrus"
	//multiple declaration
		 //var name, place string
	//multiple declaration and intialization
		 //var name, place string= "Mukunthan-2", "India-2"
	//multiple declaration and intialization with type inference
		 //var name, place = "Mukunthan-99", "India-88"
	//multiple variables with shorthand syntax
		  //name, place := "Mukunthan", "India"
		  //name, marks := "Mukunthan-re", "100"
		  //name, marks := "Mukunthan", 99
	//multiple variables with shorthand syntax and reassignment
		  //name, place := "Mukunthan", "USA"
		  //place= "UK"

	//var with brackets
		var (
			name="Mukunthan"
			place="Singapore"
		)

	//Type of variable
		fmt.Printf("topic Type is %T \n", topic)



	
	
	fmt.Println("Hello!", name, place,  topic,"<<<")
	}

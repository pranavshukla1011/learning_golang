package main

import "fmt"

func main(){
	fmt.Println("--------STUCT---------");

	/*
		STRUCT is an heterogenous data tpe, the only limitation we have on our struct is that all the fields that are to be used in a struct have to be available at runtime.
	*/

	type user struct{
		firstName string
		lastName string
		id	int
	}

	var u user

	/*
		When we initalize a struct,
		each element gets initalized to its zero value.

		Zero value for a string is "", int is 0

		so {  0}
	*/

	u.id = 1;
	u.firstName = "Pranav"
	u.lastName = "Shukla"

	fmt.Println(u)

	u2 := user{	
				id:2, 
				firstName: "Pranav", 
				lastName: "Shukla",
			  }


	fmt.Println(u2)

	/*
		We can define a struct at any level we need, 
		global or local
	*/
}
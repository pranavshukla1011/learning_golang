package main

import "fmt"

func main(){
	var firstName *string = new(string)
	fmt.Println(firstName);

	*firstName = "Pranav Shukla"

	fmt.Println(*firstName)

	// pointer arithmatic is not allowed in go.
	
	/*
		while pointer arithmatic is very powerful, it is prone to errors and altogether removed in go.
	*/

	name := "Pranav Shukla";
	pointer_name := &name;
	
	fmt.Println(pointer_name, name);

	name = "Shukla Pranav";
	
	/*
		the value pointer refers to also get changes
	*/

	fmt.Println(pointer_name, name);

}


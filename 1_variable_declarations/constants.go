package main

import "fmt"

func main(){

	fmt.Println("------------- Constants ------------");

	/*
		1. constants need to be assigned at the time of intialization

		2. value of a constant has to be determined at compile time

		3. Thus constants can be dynamically generated at compile time.
	*/

	// IMPLECITLY TYPED CONSTANT	
	/* type of a constant is decided implicitly by go*/

	const c = 3

	fmt.Println(c);

	// Thus we can do c + 3.14

	fmt.Println(c+3.14)


	// EXPLICITLY TYPED CONSTANTS
	
	const c1 int = 3;

	// Now we cannot do c1 + 3.14. Because go does not allow addition of two different types.

	//  We need to do type conversion.
	//  float32(c1) + 3.14


	fmt.Println(float32(c1) + 3.14);


}
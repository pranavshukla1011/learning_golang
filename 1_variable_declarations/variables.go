package main

import "fmt"

func main(){
	var i int
	i = 10
	fmt.Println(i);


	var f float32 = 3.14
	fmt.Println(f);

	firstName := "Pranav Shukla"
	fmt.Println(firstName)

	c := complex(1,4);
	fmt.Println(c);

	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
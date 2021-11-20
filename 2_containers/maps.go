package main

import "fmt"

func main(){
	fmt.Println("----------MAPS----------");

	m:=map[string]int{"foo":42}

	fmt.Println(m);

	fmt.Println(m["foo"])

	delete(m, "foo");

	fmt.Println(m);
}
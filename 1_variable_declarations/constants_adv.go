package main

import "fmt"

/*
	just like we have ' import block ', we can also create ' constant blocks '
*/

// const (
// 	first = iota; // iota = 0
// 	second = iota + 6 // iota = 1
// 	third = 2 << iota // iota = 3
// )

/*
	-> if we do not define the calue of a contant, inside the block,
	the contant tkae the value of the previous constant;

	-> it does not matter if we have iota as a value or not for the constant to be repeated

*/

const (
	first = iota + 6;
	second
	third
)


/*
	1!!!! --------- iota resets in every constant block ------------ !!!!

*/


const (
	forth = iota;
	fifth
	six = 3;
	seven
)

func main(){

	fmt.Println("---------- Iota and Constants Expressions ------------");

	fmt.Println(first, second, third);

	fmt.Println(forth, fifth, six, seven);

}
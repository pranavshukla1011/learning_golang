package main

import "fmt"

func main(){
    fmt.Println("-------ARRAYS------------");

	var arr [3]int;
	arr[0] = 1;
	arr[1] = 2;
	arr[2] = 3;

	fmt.Println(arr);

	fmt.Println(arr[2]);


	// arr := [3]int(4,5,6); we cannot redeclare an array

	arr2 := [3]int{4,5,6};
	fmt.Println(arr2);
}
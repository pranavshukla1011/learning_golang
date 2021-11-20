package main

import "fmt"

func main(){

	fmt.Println("------- SLICES -------");

	arr := [3]int{1,2,3}

	slice := arr[:];

	arr[1] = 33;
	slice[2] = 47;

	fmt.Println(arr);
	fmt.Println(slice)


	/*
		-> "slice" is basically pointing to the array;
		Thus,
			if we make any changes to the slice, it will be reflected in array
		and
			if we make any changes in array, it will be reflected in slice.

		->


		We usually use slices , for example when we get an array returned to us from an api and we want to use the array more freely and dynamically.

		-----> We can create a slice direclty without an ecplicitly defining an array;


		-----> slices are dynamically sized.
		As we try to change size, the array we have is going to overrun, and go is going to create a new underlying array.

		-----> 
	*/

	slice1 := []int{1,2,3};
	fmt.Println(slice1)

	slice1 = append(slice1, 56);

	fmt.Println(slice1)

	/*
		SLICE OF SLICE
	*/

	s2 := slice1[:2]
	fmt.Println(s2)

	
}
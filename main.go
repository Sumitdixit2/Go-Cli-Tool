package main

import (
	"fmt"
)

func main(){
	var num int 
	fmt.Print("Enter how many numbers to use: ")
	_,err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println("Error reading input: ",err)
		return
	}

	numbers := make([]string,num)

	for i := 0 ; i < num;i++{

		fmt.Printf("Enter number %d: " , i+1)
		fmt.Scanln(&numbers[i])
	}

	intSlice,err := parse(numbers)

	if intSlice == nil || err != nil {
		fmt.Println("Error while parsing the string slice: ",err)
	}

	Sum := Sum(intSlice)
	Average := Average(intSlice)
	Max := Max(intSlice)

	fmt.Println("Sum is: ",Sum)
	fmt.Println("Average is: ",Average)
	fmt.Println("Max is: ",Max)
}

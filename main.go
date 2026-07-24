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

	fmt.Println("numbers slice is: ",numbers)
}

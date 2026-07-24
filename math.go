package main

import (
	"fmt"

	"errors"
)

func Sum(numbers []int) int {
		sum := 0
		for _,num := range numbers{
			sum += num
		}

		return sum
}

func Average(numbers []int) (int,error){
		length := len(numbers)
		TotalSum := Sum(numbers)

		if TotalSum <= 0 {
			var err = errors.New("Total Sum is zero or less than zero")
			fmt.Printf("Error: ",err)
			return 0,err
		}

		Average := TotalSum/length

		return Average,nil
}

func Max(numbers []int) (int, error) {

		max := numbers[0]
				
		for i := 1; i < len(numbers); i++{
			if max < numbers[i]{
				max = numbers[i]
			}
		}

		return max,nil
} 

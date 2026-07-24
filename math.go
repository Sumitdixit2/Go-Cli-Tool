package main

func Sum(numbers []int) int {
		sum := 0
		for _,num := range numbers{
			sum += num
		}

		return sum
}

func Average(numbers []int) (int,err){
		length := len(numbers)
		TotalSum := Sum(numbers)

		if TotalSum <= 0 {
			fmt.Printf("Error: Total Sum is zero or less than zero")
		}

		Average := TotalSum/length

		return Average
}

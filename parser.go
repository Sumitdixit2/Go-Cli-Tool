package main

import (
	"strconv"
)

func parse(numbers []string) ([]int,error) {
		intSlice := make([]int,0,len(numbers))

		for _,str := range numbers{
			num,err := strconv.Atoi(str)

			if err != nil {

				return nil,err
	
			}	

			intSlice = append(intSlice,num)
		}

		return intSlice,nil

}

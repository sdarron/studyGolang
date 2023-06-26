// Hаписать функцию, которая переворачивает все элементы массива или
// среза местами.
// Реализовать функцию, которая в качестве аргумента принимает число n, и
// выводит в консоль все числа от 2¹ до 2^n

package main

import ( 
	"fmt"
	"math"
)

func main(){
	slice := []int {
		1,
		2,
		3,
		4,
	}
	slice = swopSlice( slice )
	fmt.Println(slice)

	n := 5
	numbersInPower( n )	
}

func numbersInPower( num int ){
	for index := 1; index <= num; index++ {		
		fmt.Printf("Число 2 в степени %d равно %d\n", index, int(math.Pow(2, float64(index))))
	}
}	


func swopSlice( slice []int ) ( []int ){
	for index := 0; index < len(slice) / 2; index++ {
		tmp := slice[index]
		slice[index] = slice[len(slice) -1 - index ]
		slice[len(slice) -1 - index] = tmp
	}
	return slice
}

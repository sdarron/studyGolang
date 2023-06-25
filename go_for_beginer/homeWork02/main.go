// Hаписать функцию, которая переворачивает все элементы массива или
// среза местами.
// Реализовать функцию, которая в качестве аргумента принимает число n, и
// выводит в консоль все числа от 2¹ до 2^n

package main

import ( 
	"fmt"
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
}

func swopSlice( slice []int ) ( []int ){
	for index := 0; index < len(slice) / 2; index++ {
		tmp := slice[index]
		slice[index] = slice[len(slice) -1 - index ]
		slice[len(slice) -1 - index] = tmp
	}
	return slice
}

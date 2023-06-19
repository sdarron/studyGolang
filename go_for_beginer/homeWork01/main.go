// Перепишите программу для подсчета площади круга, 
// используя число Пи из библиотеки math. 
// Напишите функции для подсчета площади и периметра
// прямоугольника и треугольника. 

package main

import ( 
	"fmt"
	"errors"
	"math"
)
func main()  {
	printCircleArea(2)
	printCircleArea(-2)
}


func printCircleArea( radius int ){
	circleArea, err := calculateCircleArea( radius )
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга %d см.\n", radius)	
	print("Формула для расчета круга: A=Pi*(r)2\n")
	fmt.Printf("Площать круга: %f32 см. кв. \n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным")
	}
	return float32(radius) * float32(radius) * math.Pi, nil
}
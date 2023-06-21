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
	printRectangleArea( 2, 4 )
	printTriangleArea( 8, 2)
}

func calculateTriangleArea( base int, height int ) ( int, error ){
	if base <=0 {
		return 0, errors.New("Основание треугольника не может быть отрицательным")
	}
	if height <= 0 {
		return 0, errors.New("Высота треугольника не может быть отрицательной")
	}
	return ( base * height ) / 2, nil
}

func printTriangleArea( base int, height int){
	triangleArea, err := calculateTriangleArea( base, height )
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Основание треугольнка равно %d см. и высота треугольника равна %d см.\n", base, height)	
	print("Формула для площади треугольника: S=(a * h) / 2\n")
	fmt.Printf("Площать треугольника: %d см. кв. \n", triangleArea)

}


func printRectangleArea( side1 int, side2 int ){
	rectangleArea, err := calculateRectangleArea( side1, side2)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Стороны проямоугольника равны %d и %d см.\n", side1, side2)	
	print("Формула для площади прямоугольника: S=a*b\n")
	fmt.Printf("Площать прямоугольника: %d см. кв. \n", rectangleArea)
}

func calculateRectangleArea( side1 int, side2 int ) ( int, error ){
	if side1 <=0 || side2 <= 0{
		return 0, errors.New("Сторона прямоугольнка не может быть отрицательной")
	}
	return side1 * side2, nil
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
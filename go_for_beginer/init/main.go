package main

import ( 
	"fmt"
	"errors"
)
const pi = 3.1415
func main()  {
	printCircleArea(2)
	printCircleArea(-2)
	x := 10
	p := &x
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(*&x)
	increment( &x )
	fmt.Println(x)
}

func increment(p *int){
	*p += 1
}

func printCircleArea( radius int ){
	circleArea, err := calculateCircleArea( radius )
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга %d см.\n", radius)
	
	fmt.Println("Формула для расчета круга: A=pr2\n")
	fmt.Printf("Площать круга: %f32 см. кв. \n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным")
	}
	return float32(radius) * float32(radius) * pi, nil
}
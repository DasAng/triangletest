package main

import (
	"fmt"
	"triangletest/geometry"
)

func main() {

	fmt.Println("This program will help you determine whether a triangle is Equilateral, Isosceles or Scalene")

	a,b,c := getTriangleInputs()

	tr,err := geometry.NewTriangle(a,b,c)
	if err != nil {
		panic(err)
	}

	triangleType := ""

	switch(tr.GetTriangleType()) {
	case geometry.Equilateral:
		triangleType = "Equilateral"
		break
	case geometry.Isosceles:
		triangleType = "Isosceles"
		break
	case geometry.Scalene:
		triangleType = "Scalene"
		break
	}

	fmt.Printf("Triangle is of type: %s",triangleType)

}

func getTriangleInputs() (int,int,int) {

	msg := "Enter length of side"

	a := getLengthOfTriangleSide(fmt.Sprintf("%s A: ",msg))

	b := getLengthOfTriangleSide(fmt.Sprintf("%s B: ",msg))

	c := getLengthOfTriangleSide(fmt.Sprintf("%s C: ",msg))

	fmt.Printf("A = %v\n",a)
	fmt.Printf("B = %v\n",b)
	fmt.Printf("C = %v\n",c)

	return a,b,c
}

func getLengthOfTriangleSide(msg string) (int) {

	var n int

	fmt.Print(msg)

	fmt.Scanf("%d",&n)

	return n
}

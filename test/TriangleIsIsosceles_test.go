package test

import (
	"testing"
	"triangletest/geometry"
)

func TestTriangleIsIsosceles(t *testing.T) {

	a,b,c := 11,11,10
	triangleType := geometry.Isosceles

	t.Logf("Subject: Triangle is isosceles")

	t.Logf("Given a valid triangle")

	tr,err := geometry.NewTriangle(a,b,c)
	PanicOnError(err)

	t.Logf("Then the triangle type should be %v",triangleType)
	if !tr.IsTriangleIsosceles() {
		t.Fatalf("Expected triangle type to be %v but instead was %v",triangleType,tr.GetTriangleType())
	}

	t.Logf("Test successfull")
}



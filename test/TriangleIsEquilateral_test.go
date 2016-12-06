package test

import (
	"testing"
	"triangletest/geometry"
)

func TestTriangleIsEquilateral(t *testing.T) {

	a,b,c := 11,11,11
	triangleType := geometry.Equilateral

	t.Logf("Subject: Triangle is equilateral")

	t.Logf("Given a valid triangle")

	tr,err := geometry.NewTriangle(a,b,c)
	PanicOnError(err)

	t.Logf("Then the triangle type should be %v",triangleType)
	if !tr.IsTriangleEquilateral() {
		t.Fatalf("Expected triangle type to be %v but instead was %v",triangleType,tr.GetTriangleType())
	}

	t.Logf("Test successfull")
}



package test

import (
	"testing"
	"triangletest/geometry"
)

func TestTriangleIsScalene(t *testing.T) {

	a,b,c := 11,10,12
	triangleType := geometry.Scalene

	t.Logf("Subject: Triangle is scalene")

	t.Logf("Given a valid triangle")

	tr,err := geometry.NewTriangle(a,b,c)
	PanicOnError(err)

	t.Logf("Then the triangle type should be %v",triangleType)
	if !tr.IsTriangleScalene() {
		t.Fatalf("Expected triangle type to be %v but instead was %v",triangleType,tr.GetTriangleType())
	}

	t.Logf("Test successfull")
}



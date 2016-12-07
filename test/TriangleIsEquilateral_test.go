package test

import (
	"testing"
	"triangletest/geometry"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTriangleIsEquilateral(t *testing.T) {

	a,b,c := 11,11,11
	Convey("Scenario 1: 3 equal length sides should return equilateral triangle type",t, func() {

		Convey(fmt.Sprintf("Given lengths a = %v, b = %v , c = %v",a,b,c),func() {

			Convey("And should form a valid triangle",func() {

				tr,err := geometry.NewTriangle(a,b,c)

				So(err,ShouldBeNil)

				Convey("When a,b,c are all equal", func() {

					Convey("Then the triangle type should be equilateral",func() {
						So(tr.IsTriangleEquilateral(),ShouldBeTrue)
					})
				})
			})
		})
	})
}



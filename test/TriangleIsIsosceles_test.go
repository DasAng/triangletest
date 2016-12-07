package test

import (
	"testing"
	"triangletest/geometry"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTriangleIsIsosceles(t *testing.T) {

	a,b,c := 11,11,10

	Convey("Scenario 1: 2 equal length sides should return isosceles triangle type",t, func() {

		Convey(fmt.Sprintf("Given lengths a = %v, b = %v , c = %v",a,b,c),func() {

			Convey("And should form a valid triangle",func() {

				tr,err := geometry.NewTriangle(a,b,c)

				So(err,ShouldBeNil)

				Convey("When a and b are equal", func() {

					Convey("Then the triangle type should be isosceles",func() {
						So(tr.IsTriangleIsosceles(),ShouldBeTrue)
					})
				})
			})
		})
	})
}



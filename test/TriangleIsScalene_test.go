package test

import (
	"testing"
	"triangletest/geometry"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestTriangleIsScalene(t *testing.T) {

	a,b,c := 11,10,12

	Convey("Scenario 3: no equal length sides should return scalene triangle type",t, func() {

		Convey(fmt.Sprintf("Given lengths a = %v, b = %v , c = %v",a,b,c),func() {

			Convey("And should form a valid triangle",func() {

				tr,err := geometry.NewTriangle(a,b,c)

				So(err,ShouldBeNil)

				Convey("When a,b and c are not equal", func() {

					Convey("Then the triangle type should be scalene",func() {
						So(tr.IsTriangleScalene(),ShouldBeTrue)
					})
				})
			})
		})
	})
}



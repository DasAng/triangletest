package geometry

import (
	"sort"
	"fmt"
)

const (
	Equilateral = iota
	Isosceles = iota
	Scalene = iota
)

// A triangle has 3 sides a,b and c.
// In addition the field triangleType is the triangle type and can be one of the following values:
// Equilateral = 0, Isosceles = 1 or Scalene = 2.
type triangle struct {
	a    int
	b    int
	c    int
	triangleType int
}

// NewTriangle returns a new triangle instance
// or an error if the specified sides a,b,c does not form a valid triangle.
// If an error occurs then the returned instance is nil.
func NewTriangle(a int, b int, c int) (*triangle, error) {

	if !isTriangleValid(a, b, c) {
		return nil, fmt.Errorf("Sides a,b,c does not form a valid triangle")
	}

	t := &triangle{a:a, b:b, c:c}
	t.determineTriangleType()

	return t,nil
}

// isTriangleValid returns true if the specified sides a,b and c forms a valid triangle otherwise false.
func isTriangleValid(a int, b int, c int) bool {

	sortedNumbersAsc := sortNumbersAscending([]int{a, b, c})
	return performInequalityMethodFast(sortedNumbersAsc[0], sortedNumbersAsc[1], sortedNumbersAsc[2])
}

// sortNumbersAscending sorts the numbers n in ascending order.
func sortNumbersAscending(n []int) []int {

	numbers := sort.IntSlice(n)
	sort.Sort(numbers)
	return numbers
}

// performInequalityMethodFast is using the Inequality Theorem to prove that all three side lengths
// forms a triangle. If the sum of the 2 smallest sides a and b is greater than the largest side c,
// then the sides forms a valid triangle.
// Negative numbers is not a problem since it will always fail the Inequality test.
func performInequalityMethodFast(a int, b int, c int) bool {
	return a + b > c
}

// areAllSidesEqual returns true if all sides of a triangle are equal
func (m *triangle) areAllSidesEqual() bool {

	return m.a == m.b && m.a == m.c
}

// areTwoSidesEqual returns true if two sides of a triangle are equal
func (m *triangle) areTwoSidesEqual() bool {

	return !m.areAllSidesEqual() && !m.areNoSidesEqual()
}

// areNoSidesEqual returns true if no sides of a triangle are equal
func (m *triangle) areNoSidesEqual() bool {

	return m.a != m.b && m.b != m.c && m.c != m.a
}

// determineTriangleType determines the triangle type and save this information in the triangleType field.
func (m *triangle) determineTriangleType() {

	if m.areAllSidesEqual() {
		m.triangleType = Equilateral
	} else if m.areTwoSidesEqual() {
		m.triangleType = Isosceles
	} else {
		m.triangleType = Scalene
	}
}

// IsTriangleEquilateral returns true if the triangle is equilateral, i.e all sides are equal
func (m *triangle) IsTriangleEquilateral() bool {
	return m.triangleType == Equilateral
}

// IsTriangleIsosceles returns true if the triangle is Isosceles, i.e two sides are equal
func (m *triangle) IsTriangleIsosceles() bool {
	return m.triangleType == Isosceles
}

// IsTriangleScalene returns true if the triangle is Scalene, i.e no sides are equal
func (m *triangle) IsTriangleScalene() bool {
	return m.triangleType == Scalene
}

// GetAllSides returns all the sides of the triangle as an array
func (m *triangle) GetAllSides() ([]int) {
	return []int{m.a,m.b,m.c}
}

// GetTriangleType returns an enumeration value where:
// Equilateral = 0, Isosceles = 1, Scalene = 2
func (m *triangle) GetTriangleType() int {
	return m.triangleType
}





# INTRODUCTION #

This is a program that allows you to determine whether a triangle is one of the following types:

- Equilateral
- Isosceles
- Scalene

Given the length of 3 sides of a triangle this program will tell you whether this forms a valid triangle and what type of triangle it is.

# USAGE #

To run the application type this in a terminal

**Linux**
 
 ````bash
 triangletest
 ````
 
 **Windows**
 
 ````bash
 triangletest.exe
 ````
 
 You will be prompted to enter the length of the 3 sides of a triangle.
 
 The program will then display the result.
 
# BUILD FROM SOURCE #
 
 In order to build the program from source you will need to have *Go* installed:
 
 Visit https://golang.org/dl/ and download the appropriate binary for your platform
 
### Compile the source ###
 
 Before compiling your application make sure you have set up the project structure correctly. 
 Create a directory, eg. lets called it workspace/src
 
 ````bash
 sudo mkdir -p /workspace/src
 cd /workspace/src
````

Now git clone the project inside **/workspace/src**

Set up *GOPATH* environment variable to point to /workspace

*Linux*

````bash
export GOPATH=/workspace
````

*Windows*

````bash
set GOPATH=/workspace
````

We are now ready to compile the application

````bash
go get -v -t
go build
````

You should now have a compiled binary of your application called **triangletest** in your project folder

# RUN TEST WITH CODE COVERAGE #

To run all the test with code coverage report:

````bash
go test -v triangletest/test -coverprofile=code.cover.out -coverpkg=./...
````

Then view the report in a nice HTML by typing this:

````bash
go tool cover -html=code.cover.out
````

# UNDERSTANDING THE CODE #

The scope of this program is to determine whether a triangle is either equilateral, isosceles or scalene.

We are not concerned with the triangle's angle so basically we don't care whether the triangle is:

- Equiangular
- Obtuse
- Right
- Acute

### Testing ###

The first obvious thing that spring to mind is how can we tests this?

Using **TDD** we can make sure that we have defined our test scenarios and what we wants to test first.
Then we can start fleshing out the logic to make the test pass.

Basically we can start out by setting the expectation of what we want to achieve.

Using **BDD** approach we can communicate this to others to understand our domain specific problem:

````gherkin
Story: Returns if triangle is either equilateral, isosceles or scalene
As a user
In order to understand triangle type 
I want to know if 3 sides of a triangle forms an equilateral, isosceles or scalene type
````

````gherkin
Scenario 1: 3 equal length sides should return equilateral triangle type
Given lengths a = 10, b = 10 , c = 10
And should form a valid triangle
When a,b,c are all equal
Then the triangle type should be equilateral
````

````gherkin
Scenario 2: 2 equal length sides should return isosceles triangle type
Given lengths a = 10, b = 10 , c = 5
And should form a valid triangle
When a and b are equal
Then the triangle type should be isosceles
````

````gherkin
Scenario 3: no equal length sides should return scalene triangle type
Given lengths a = 10, b = 11 , c = 12
And should form a valid triangle
When a,b and c are not equal
Then the triangle type should be scalene
````

I want to have as minimal a dependency on external libraries and packages in order to keep the binary size small and easy to understand.

We will only be using the package http://goconvey.co/ to help us define the BDD specs in our tests.

#### The Inequality Theorem ####

In order to validate whether 3 lengths can form a triangle we can use the Inequality theorem that says that given three sides a,b and c:

- a + b > c
- b + c > a
- a + c > b

All three conditions must be true to meet the requirement.

Looking at those three conditions we can see that we don't need to perform all three conditions one by one.

If we can eliminate any one of those three rules then we can stop immediately and conclude that the triangle is invalid.

Basically we can state our conditions a bit differently:

`The sum of two smallest number should be greater than the largest number`

This condition will meet the requirement stated by the Inequality theorem.

Instead of testing each condition one by one as stated in the Inequality theorem we can figure out which number is the largest one and add the other 2 numbers together.

We can do this in Go by sorting our 3 lengths in ascending order and adding the first two items together and see if it is larger than the last item.

#### Equilateral test ####

A triangle is equilateral if the length of all three sides are equal.

This can easily be determined using if conditions in Go:

`if a == b && b == c`

#### Scalene test ####

A triangle is scalene if the length of all three sides are not at all equal.

This can easily be determined using if conditions in Go:

`if a != b && b != c && c != a`

#### Isosceles test ####

A triangle is isosceles if the length of two sides are equal.

We can do this by making sure it is both not equilateral and not scalene.

### Program logic ###

Given the knowledge we have and the test scenarios we wish to perform we can then implement our business logic.

We want to allow the user to enter 3 lengths at which we will want to validate whether the 3 lengths can form a valid triangle.

If the requirement is not met we will inform the user that he/she has entered an invalid triangle.

If the triangle is valid we will proceed by determining the triangle type using the Equilateral, Isosceles and Scalene checks as mentioned above.

The output will the string: "The entered triangle is a X triangle" where X can either be "Equilateral","Isosceles" or "Scalene"

### Implementation ###

#### The triangle struct ####

We will need to have a *triangle* struct which will contain the lengths of the three sides we call them a,b and c

````go
type triangle struct {
        a int
        b int
        c int
}
````

The reason we have chosen to make the triangle struct private is because we don't want to allow this triangle to be instantiated with invalid sides.

We will instead expose a factory method that allow us to instantiate a new triangle given 3 lengths. See code snippet below:

````go
func NewTriangle(a int, b int, c int) (*triangle, error) {

	if !isTriangleValid(a, b, c) {
		return nil, fmt.Errorf("Sides a,b,c does not form a valid triangle")
	}

	t := &triangle{a:a, b:b, c:c}
	t.determineTriangleType()

	return t,nil
}
````

As you can see this method will only return a valid triangle if it passes the validation.

Furthermore we also wanted to determine the triangle type immediately so that we can store this information in the triangle object returned.

This way we do not need to determine the triangle type every time we ask for that information.

The rest of the code is pretty straightforward.
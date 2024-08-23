package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is the initialization")
}

func main() {
	//x := rand.Intn(351)

	// if x > 0 && x <= 100 {
	// 	fmt.Printf("Value %v is between 0 and 100", x)
	// }

	// if x > 100 && x <= 200 {
	// 	fmt.Printf("Value %v is between 100 and 200", x)
	// }

	// if x > 200 && x <= 250 {
	// 	fmt.Printf("Value %v is between 200 and 250", x)
	// }

	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))
	// fmt.Printf("\n\t rand number %v\n", rand.Intn(3))

	// switch {
	// case x > 0 && x < 100:
	// 	fmt.Printf("Value %v is between 0 and 100", x)
	// case x > 100 && x < 200:
	// 	fmt.Printf("Value %v is between 0 and 100", x)
	// case x > 200 && x < 250:
	// 	fmt.Printf("Value %v is between 200 and 250", x)
	// default:
	// 	fmt.Println("DEFAULT case")
	// }

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Value of x : %v\n", x)
	fmt.Printf("Value of y : %v\n", y)

	// if x < 4 && y < 4 {
	// 	fmt.Printf("x[%v] and y[%v] both are less than 4\n", x, y)
	// } else if x > 6 && y > 6 {
	// 	fmt.Printf("x[%v] and y[%v] are both greater than 6\n", x, y)
	// } else if x >= 4 && x <= 6 {
	// 	fmt.Printf("x[%v] is less than or equal to 4 and less than or equal to 6\n", x)
	// } else if x != 5 {
	// 	fmt.Printf("x[%v] is not 5\n", x)
	// } else {
	// 	fmt.Println("None of the previous test cases were met")
	// }

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x[%v] and y[%v] both are less than 4\n", x, y)
	case x > 6 && y > 6:
		fmt.Printf("x[%v] and y[%v] are both greater than 6\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("x[%v] is less than or equal to 4 and less than or equal to 6\n", x)
	case y != 5:
		fmt.Printf("y[%v] is not 5\n", x)
	default:
		fmt.Println("None of the previous test cases were met")
	}
}

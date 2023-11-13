package main

import (
	"fmt"
	"math"
)

// func main() {
// 	// for i := 1; i <= 20; i++ {
// 	// 	if i%3 == 0 && i%5 == 0 {
// 	// 		fmt.Println(i, "FizzBuzz")
// 	// 	} else if i%3 == 0 {
// 	// 		fmt.Println(i, "Fizz")
// 	// 	} else if i%5 == 0 {
// 	// 		fmt.Println(i, "Buzz")
// 	// 	} else {
// 	// 		fmt.Println(i)
// 	// 	}
// 	// }

// 	// for i := range 20 {
// 	// 	if i%3 == 0 && i%5 == 0 {
// 	// 		fmt.Println(i, "FizzBuzz")
// 	// 	} else if i%3 == 0 {
// 	// 		fmt.Println(i, "Fizz")
// 	// 	} else if i%5 == 0 {
// 	// 		fmt.Println(i, "Buzz")
// 	// 	} else {
// 	// 		fmt.Println(i)
// 	// 	}
// 	// }


// 		for i := 0; i <= 20; i++ {
// 			if (math.Mod(float64(i), 3) == 0 && math.Mod(float64(i), 5) == 0) {
// 				fmt.Println(i, "FizzBuzz")
// 			}
// 			else if (math.Mod(float64(i), 3) == 0) {
// 				fmt.Println(i, "Fizz")
// 			} else if (math.Mod(float64(i), 5) == 0) {
// 				fmt.Println(float64(i), "Buzz")
// 			} else {
// 				fmt.Println(i)
// 			}
// 		} 
// }


func main() {
	// for i := 0; i <= 20; i++ {
	// 	if math.Mod(float64(i), 3) == 0 && math.Mod(float64(i), 5) == 0 {
	// 		fmt.Println(i, "FizzBuzz")
	// 	} else if math.Mod(float64(i), 3) == 0 {
	// 		fmt.Println(i, "Fizz")
	// 	} else if math.Mod(float64(i), 5) == 0 {
	// 		fmt.Println(i, "Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }
	for i := 0; i <= 20; i++ {
		fmt.Println(printBuzz(i))
	}
}

func printBuzz(num int) string {
	if math.Mod(float64(num), 3) == 0 && math.Mod(float64(num), 5) == 0 {
		return fmt.Sprintf("%v FizzBuzz", num)
	} else if math.Mod(float64(num), 3) == 0 {
		return fmt.Sprintf("%v Fizz", num)
	} else if math.Mod(float64(num), 5) == 0 {
		return fmt.Sprintf("%v Buzz", num)
	} else {
		return fmt.Sprintf("%v", num)
	}
}



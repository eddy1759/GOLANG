package main

import (
	"fmt"
)

// func main() {
// 	var firstname string = "Edet"
// 	var lastname string
// 	lastname = "Asuquo"
// 	var age int
// 	age = 29
// 	stack := "golang"
// 	const pi = 3.14159

// 	fmt.Println("Hello, my name is "+ firstname + " " + lastname + " and i am a golang developer " + " and i am ", age, " years old")
// 	fmt.Println()
// 	fmt.Println("Hey there, my name is", firstname, lastname + " a", age, " years old golang developer")
// 	fmt.Println()
// 	fmt.Println("Hey there, my name is", firstname, lastname + " a", age, " years old golang developer and i am learning", stack, "stack")
// 	fmt.Println()
// 	fmt.Println("The value of pi is: ", pi)
// 	//type check the variables
// 	fmt.Printf("pi is of type %T\n", pi)

// 	num1 := 5
// 	num2 := 10

// 	var num3 float64 = 15
// 	var num4 float64 = 25.0

// 	sum1 := num1 + num2
// 	sum2 := num3 + num4
// 	subtract1 := num2 - num1
// 	subtract2 := num3 - num4
// 	multiply1 := num1 * num2
// 	multiply2 := num3 * num4
// 	divide1 := num2 / num1
// 	divide2 := num4 / num3
// 	remainder1 := num2 % num1
// 	exponent1 := num1 ^ num2

// 	fmt.Println("The sum of", num1, "and", num2, "is", sum1)
// 	fmt.Println("The sum of", num3, "and", num4, "is", sum2)
// 	fmt.Println("The difference between", num2, "and", num1, "is", subtract1)
// 	fmt.Println("The difference between", num3, "and", num4, "is", subtract2)
// 	fmt.Println("The product of", num1, "and", num2, "is", multiply1)
// 	fmt.Println("The product of", num3, "and", num4, "is", multiply2)
// 	fmt.Println("The quotient of", num2, "and", num1, "is", divide1)
// 	fmt.Println("The quotient of", num4, "and", num3, "is", divide2)
// 	fmt.Println("The remainder of", num2, "and", num1, "is", remainder1)
// 	fmt.Println("The exponent of", num1, "and", num2, "is", exponent1)
// }

// func main() {
// 	// age := 29

// 	// if age > 29 {
// 	// 	fmt.Println("You're eligible to vote")
// 	// } else {
// 	// 	fmt.Println("You're not eligible to vote")
// 	// }
// 	// for i := 1; i <= 10; i++ {
// 	// 	fmt.Println("Interation: ", i)
// 	// }
// 	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
// 	// switch days[2] {
// 	// case "Monday":
// 	// 	fmt.Println("Today is Monday")
// 	// case "Tuesday":
// 	// 	fmt.Println("Today is Tuesday")
// 	// case "Wednesday":
// 	// 	fmt.Println("Today is Wednesday")
// 	// case "Friday":
// 	// 	fmt.Println("Today is Friday")
// 	// default:
// 	// 	fmt.Println("It is another day")
// 	// }

// 	a := []int{1, 2, 3, 4, 5} 
// 	fmt.Println(a) // [1 2 3 4 5]
// 	a[2] = 10-2 // 8
// 	fmt.Println(a) // [1 2 8 4 5]
// 	a = append(a, 6) 
// 	fmt.Println(a) // [1 2 8 4 5 6]
// 	fmt.Println(a[2:4]) // [8 4]
// 	fmt.Println(a[2:]) // [8 4 5 6]
// 	fmt.Println(a[:4]) // [1 2 8 4]
// 	fmt.Println(a[:]) // [1 2 8 4 5 6]
// 	fmt.Println(a[1:4]) // [2 8 4]
// 	a = append(a[:2], a[3:]...) // [1 2 4 5 6]
// 	fmt.Println(a)
// 	fmt.Println(len(a)) // 5
// 	fmt.Println("cap: ", cap(a)) // 10
// 	b := make([]int, 10)
// 	fmt.Println(b)
// 	fmt.Println(len(b))
// 	fmt.Println(cap(b))
// 	a = append(a, b...)
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))
// }.

// func main() {
// 	var name string 
// 	var age int
// 	fmt.Println("Enter your name: ")
// 	fmt.Scanln(&name)
// 	fmt.Println("Enter your age: ")
// 	fmt.Scanln(&age)
// 	fmt.Println("Hello", name, " you're ", age, " years old")
// }

func main () {
	var x int = 10
	var y  = &x
	fmt.Println("Value stored in x: ", x)
	fmt.Println("Address of x: ", &x)
	fmt.Println("Value stored in y: ", y)
	fmt.Println("Address of y: ", &y)
	fmt.Println("Value pointed to by y: ", *y)
	fmt.Println("Value pointed to by y: ", *(&x))
	fmt.Println("Value pointed to by y: ", *&y)
	*y = 20
	fmt.Println(("value stored stored in  x(*y) after changing *y: "), x)
	fmt.Println((x))
}
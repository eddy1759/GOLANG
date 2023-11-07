package main

import "fmt"

func areaOfRectangle(length, breadth float64) float64 {
    return length * breadth
}

func areaOfCircle(radius float64) float64 {
    return 3.14 * radius * radius
}

func main() {
    fmt.Println("What area would you like to calculate?")
    fmt.Println("1. Rectangle")
    fmt.Println("2. Circle")

    var choice int
    _, err := fmt.Scanf("%d", &choice)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid choice.")
        return
    }

    // Consume the newline character left in the input buffer
    fmt.Scanln()

    switch choice {
    case 1:
        var length, breadth, area float64
        fmt.Println("Enter the length of the rectangle:")
        _, err := fmt.Scanf("%f", &length)
        if err != nil {
            fmt.Println("Invalid input for length. Please enter a valid number.")
            return
        }
		fmt.Scanln()
		fmt.Println("Enter the breadth of the rectangle:")
        _, err = fmt.Scanf("%f", &breadth)
		fmt.Scanln()
        if err != nil {
            fmt.Println("Invalid input for breadth. Please enter a valid number.")
            return
        }
        area = areaOfRectangle(length, breadth)

        if area > 10 {
            fmt.Println("The area of the rectangle is greater than 10:", area)
        } else {
            fmt.Println("The area of the rectangle is less than 10:", area)
        }

    case 2:
        var radius, area float64
        fmt.Println("Enter the radius of the circle:")
        _, err := fmt.Scanf("%f", &radius)
		fmt.Scanln()
        if err != nil {
            fmt.Println("Invalid input for radius. Please enter a valid number.")
            return
        }
        area = areaOfCircle(radius)

        if area > 10 {
            fmt.Println("The area of the circle is greater than 10:", area)
        } else {
            fmt.Println("The area of the circle is less than 10:", area)
        }

    default:
        fmt.Println("Invalid input. Please enter a valid choice.")
		return
    }
}

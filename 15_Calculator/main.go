package main

import "fmt" //Println()  Scanln()

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

/*
func main() {
	var choice int
	var a, b int
	fmt.Println(" enter 1--> Addition\n 2--> subtraction\n 3--> product\n 4--> division 5--> exit")
	fmt.Scanln(&choice)

	fmt.Println("Enter first digit: ")
	fmt.Scanln(&a)

	fmt.Println("Enter second digit: ")
	fmt.Scanln(&b)

	if choice == 1 {
		addition := add(a, b)
		fmt.Println("sum = ", addition)

	} else if choice == 2 {
		diff := sub(a, b)
		fmt.Println("subtraction = ", diff)

	} else if choice == 3 {
		prod := product(a, b)
		fmt.Println("product = ", prod)

	} else if choice == 4 {
		division := div(a, b)
		fmt.Println("quotient = ", division)

	} else {
		fmt.Println("Enter valid choice")
	}
}*/

func main() {
	var choice int
	var a, b int
	for {
		fmt.Println("\nenter operation \n 1--> Addition\n 2--> subtraction\n 3--> product\n 4--> division \n 5--> exit")
		fmt.Scanln(&choice)

		if choice <= 0 || choice > 5 {
			fmt.Println("Enter choice between 1 - 5")
			continue
		}

		if choice == 5 {
			break
		}

		fmt.Println("Enter first digit: ")
		fmt.Scanln(&a)

		fmt.Println("Enter second digit: ")
		fmt.Scanln(&b)

		switch choice {
		case 1:
			addition := add(a, b)
			fmt.Println("sum = ", addition)
		case 2:
			diff := sub(a, b)
			fmt.Println("subtraction = ", diff)
		case 3:
			prod := product(a, b)
			fmt.Println("product = ", prod)
		default:
			division := div(a, b)
			fmt.Println("quotient = ", division)
		}
	}
}

package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

// 1- Create a division function (div) with 2 integer parameters and a float return
func div(a int, b int) float64 {
	return float64(a) / float64(b)
}

func goCal(firstNum int, secondNum int) {
	fmt.Println(firstNum, "+", secondNum, "=", plus(firstNum, secondNum))
	fmt.Println(firstNum, "-", secondNum, "=", minus(firstNum, secondNum))
	fmt.Println(firstNum, "*", secondNum, "=", multi(firstNum, secondNum))

	if secondNum == 0 {
		fmt.Println("Not divisible by zero.")
	} else {
		fmt.Println(firstNum, "/", secondNum, "=", div(firstNum, secondNum))
	}
}

func main() {
	var firstNum int
	var secondNum int
	var name string

	fmt.Print("\nHello, What's your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("\nWelcome, %s! Let's do some basic calculations.\n\n", name)

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &firstNum)

	// 2- Ask the user to enter the second number and store the value in "secondNum"
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &secondNum)

	// 3- Call the "goCal" function with the proper parameters
	goCal(firstNum, secondNum)
}

// this is the calculator code written in go lang
package main

import "fmt"

// this is the function for handling the addition 
func add(x , y int) int {
	return x+y
}

// this is the function for handling substraction
func substraction( x, y int) int {
	return x-y
}
// this is the fucntion for handling the  multiplication
func product(x,y int) int {
	return x*y
}
// this is the function for handling the division
func division(x,y int) int {
	return x/y
}
//  this is the main method
func main() {
	
	for{
		fmt.Print("enter 1-> addition 2-> substraction ")
		fmt.Println("enter 3-> Product 4->  Quotient 5 -> Quit")
	
	var choice , num1 , num2 int 
	fmt.Scanln(&choice)

	if choice == 5{
		break;
	}

	fmt.Println("enter Number 1")
	fmt.Scanln(&num1)
	fmt.Println("enter Number 2")
	fmt.Scanln(&num2)
	
	// checking for  the conditions
	if (choice == 1) {
		fmt.Println("Performing Addition ")
		sum := add(num1 , num2)
		fmt.Println("Sum = " , sum)
	}else if (choice == 2) {
		fmt.Println("Performing  Substraction")
		sub := substraction(num1,num2)
		fmt.Println("Substraction = " , sub)
	}else if choice ==3{
		fmt.Println("Performing Multiplication")
		mul := product(num1,num2)
		fmt.Println("Product = " , mul)
	}else if choice == 4 {
		fmt.Println("Performing  Division")
		div := division(num1,num2)
		fmt.Println("Quotient = " , div)
	}else {
		fmt.Println("incorrect choice")
	}
}
}
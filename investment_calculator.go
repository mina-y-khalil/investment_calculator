package main

import (
	"math" // we need math package for power function there is no built-in power operator in Go
	"fmt" // this package is used for printing output to the console this also get import from terminal
) 

func main () {
	const inflationRate = 2.5
	var investmentAmount float64  // this value will be overridden by user input that is why we are not assigning any value here
	expectedReturnRate := 5.5
	var years float64


	// user inputs
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount) // pointers to read input from user (single word input)

	fmt.Print("Expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate) // pointers to read input from user

	fmt.Print("Years to invest: ")
	fmt.Scan(&years) // pointers to read input from user

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate /100) , years)
	futureRealValue := futureValue /  math.Pow((1 + inflationRate /100) , years)

	fmt.Println("Future Value of Investment: ", futureValue) // the ln after println is to print the output to the console in new line
	fmt.Println("Future Real Value of Investment (adjusted for inflation): ", futureRealValue)
}  
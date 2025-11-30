package main

import (
	"math" // we need math package for power function there is no built-in power operator in Go
	"fmt" // this package is used for printing output to the console
) 

func main () {
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5 // we removed the var and replaced it with := for type inference we can do this only when we are initializing the variable without using var keyword

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate /100) , years)
	fmt.Println("Future Value of Investment: ", futureValue) // the ln after println is to print the output to the console in new line
}  
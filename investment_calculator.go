package main

import (
	"math" // we need math package for power function there is no built-in power operator in Go
	"fmt" // this package is used for printing output to the console
) 

func main () {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow((1 + expectedReturnRate /100) , years)
	fmt.Println("Future Value of Investment: ", futureValue) // the ln after println is to print the output to the console in new line
}  
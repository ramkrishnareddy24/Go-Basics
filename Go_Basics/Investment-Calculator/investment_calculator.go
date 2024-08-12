package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnedRate = 5.5

	// var investmentAmount float64 = 1000
	// var investmentAmount, years float64 = 1000, 10
	// var expectedReturnedRate = 5.5
	// investmentAmount, years, expectedReturnedRate := 1000.0, 10.0, 5.5
	// var years float64 = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnedRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	investmentAmount = 2000
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnedRate, years)

	// fmt.Printf("Future Value: %v\n", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.0f\n", futureValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n",futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV,formattedRFV)

	// fmt.Printf(`Future Value: %.1f\n
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)

	outputText("Output Text")
}

func outputText(text1 string) {
	fmt.Print(text1)
}

func calculateFutureValue(investmentAmount, expectedReturnedRate, years float64) ( float64,  float64) {
	// fv = investmentAmount * math.Pow(1+expectedReturnedRate/100, years)
	// rfv = fv / math.Pow(1+inflationRate/100, years)
	//return
	fv := investmentAmount * math.Pow(1+expectedReturnedRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	 
}

package main

import "fmt"

type floatMap map[string]float64

func (fm floatMap) output(){
	fmt.Println(fm)
}

func main(){
	userNames := make([]string,2,5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 4.6
	courseRatings["React"] = 4.7
	courseRatings["Angular"] = 4.7

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index: ",index)
		fmt.Println("Value: ",value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ",key)
		fmt.Println("Value: ",value)
	}
}
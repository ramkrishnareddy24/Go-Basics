package functionsarevalues

import (
	"fmt"
)

type transformFun func (int) int 

func main(){
	numbers := []int{1,2,3,4}
	moreNumbers := []int{5,1,2,3,4}
	doubled := transformNumbers(&numbers,double)
	tripled := transformNumbers(&numbers,triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transF1 := getTransformerFunction(&numbers)
	transF2 := getTransformerFunction(&moreNumbers)

	// Anonymous function
	transN := transformNumbers(&numbers,transF1)


	moreTrans := transformNumbers(&moreNumbers,transF2)

	fmt.Println(transN)
	fmt.Println(moreTrans)

	// double := createTransformer(2)
	// triple := createTransformer(3)
	// fmt.Println(double)
	// fmt.Println(triple)
}

func getTransformerFunction(numbers *[]int) transformFun {
	if(*numbers)[0]==1{
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformFun) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number*2;
}

func triple(number int) int {
	return number*3;
}

//closure func
func createTransformer(factor int) func(int) int {
	return func (number int) int {
		return number * factor
	}
}
package lists

import (
	"fmt"
)

// type Product struct {
// 	id    string
// 	title string
// 	price float64
// }

func main() {
	prices := []float64{10.99,9.99,8.99,7.99}
	// fmt.Println(prices[0:1])
	prices[1] = 5.99
	// fmt.Println(prices[0:2])

	upadtedPrices := append(prices,2.99)
	fmt.Println(upadtedPrices,prices)
	// exercise()

	discountPrices := []float64{101.99,99.99,20.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func exercise() {
// 	//1.
// 	hobbies := []string{"Read", "Watch", "Listen"}
// 	fmt.Println(hobbies)

// 	//2.
// 	fmt.Println(hobbies[0])
// 	featuredHobbies := hobbies[1:]
// 	fmt.Println(featuredHobbies)

// 	//3.
// 	featuredHobbiesB := hobbies[:2]
// 	fmt.Println(featuredHobbiesB)

// 	//4.
// 	featuredHobbiesC := append(featuredHobbiesB, hobbies[2])
// 	featuredHobbiesD := featuredHobbiesC[1:]
// 	fmt.Println(featuredHobbiesD)

// 	//5.
// 	courseGoals := []string{"Golang", "Docker"}
// 	fmt.Println(courseGoals)

// 	//6.
// 	courseGoals[1] = "Gin"
// 	courseGoals = append(courseGoals, "React")
// 	fmt.Println(courseGoals)

// 	//7.
// 	products := []Product{{"001", "FirstProduct", 999.99}, {"002", "SecondProduct", 888.99}}
// 	products = append(products, Product{"003", "ThirdProduct", 777.99})
// 	fmt.Println(products)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	fmt.Println(productNames)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	prices := [4]float64{10.99, 9.99, 8.99, 7.99}
// 	fmt.Println(prices)
// 	fmt.Println(prices[2])

// 	featuredPriceA := prices[1:3]
// 	fmt.Println("Slices:", featuredPriceA)

// 	featuredPriceB := prices[:3]
// 	fmt.Println("Slices:", featuredPriceB)

// 	featuredPriceC := prices[1:]
// 	fmt.Println("Slices:", featuredPriceC)

// 	fmt.Println(len(featuredPriceC),cap(featuredPriceC))

// }

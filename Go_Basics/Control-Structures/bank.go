package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const passbook = "passbook.txt"

func main() {
	var accountBalance, error = fileops.GetFloatFromFile(passbook)
	if error != nil {
		fmt.Println("ERROR")
		fmt.Println(error)
		fmt.Println("-----")
		panic("Can't continue Sorry")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 ",randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Println("Select Choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your Choice: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance: ", accountBalance)
		case 2:
			fmt.Println("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}

			accountBalance = accountBalance + depositAmount
			fmt.Println("current account balance after deposit: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, passbook)
		case 3:
			fmt.Println("Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient Funds.")
				continue
			}

			accountBalance = accountBalance - withdrawAmount
			fmt.Println("current account balance after withdraw: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, passbook)
		default:
			fmt.Println("Thanks for choosing our bank!")
			fmt.Println("Goodbye!")
			return
		}
	}

}

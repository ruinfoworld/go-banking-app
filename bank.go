package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"
func main(){
	var accountBalance, err  = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil{
		fmt.Println("Error : ")
		fmt.Println(err)
		fmt.Println("===============")
		// panic("Can't continue, sorry!") 
	}
	fmt.Println("Welcome to the Go Banking System")
	fmt.Println("Reach us 24/7  at :", randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice	== 0
		// if wantsCheckBalance {

		switch choice {
			case 1 : 
				fmt.Println("Your Balance is", accountBalance)
			case 2 : 
				var depositeAmount float64
				fmt.Print("Please enter deposite amount: ")
				fmt.Scan(&depositeAmount)

				if depositeAmount <= 0 {
					fmt.Println("Invalid amount. Amount must be greater than zero.")
					continue
				}

				accountBalance += depositeAmount
				fmt.Println("Your new balance is: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			case 3 : 
				var withdrowAmaount float64
				fmt.Print("Please enter withdraw amount: ")
				fmt.Scan(&withdrowAmaount)

				if withdrowAmaount <= 0 {
					fmt.Println("Invalid amount. Amount must be greater than zero.")
					continue
				}

				if withdrowAmaount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have")
					continue
				}
				accountBalance -= withdrowAmaount
				fmt.Println("Your new balance is: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thank you for choosing our bank")
				return
		}
	}
}
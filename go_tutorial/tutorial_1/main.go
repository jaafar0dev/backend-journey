package main

import (
	"fmt"
	"errors"
	"time"
)

// Two sum problem, return indices of the two numbers such that they add up to a specific target.

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// ATM Simulator
func atm(balance int) {
	
	a := 0
	mainBalance := balance
	var balanceError error
	var transHis []string
	
	for a < 1 {
		
		fmt.Println("\nYour balance is $", mainBalance)
		fmt.Println("What would you like to do?")
		fmt.Println("1.Withdraw")
		fmt.Println("2.Transfer Money")
		fmt.Println("3.Check Balance")
		fmt.Println("4.Transaction History")
		fmt.Println("5.Exit")

		var input int16
		fmt.Scan(&input)
		
		if input == 1 {
			fmt.Print("How much would you like to withdraw? ")
			var amount int
			fmt.Scan(&amount)
			if amount > mainBalance {
				balanceError = errors.New("Your balance is too low for this transaction")
			} else {
				mainBalance -= amount
				fmt.Print("Withdraw successful for amount $", amount, "\n")
				
				metaData := "You withdrew " + string(amount) + " at " + string(time.Now()) 
				append(transHis, metaData)
			}
			if balanceError!=nil{
				fmt.Println(balanceError.Error())
			}
			
		}
		
		if input == 2 {
			fmt.Print("What account number do you want to transfer to ")
			var accountNumber int
			fmt.Scan(&accountNumber)
			
			fmt.Print("How much do you want to transfer? ")
			var amount int
			fmt.Scan(&amount)
			
			if amount > mainBalance {
				balanceError = errors.New("Your balance is too low for this transaction")
				
			} else {
				mainBalance = mainBalance - amount
				fmt.Println("Transfer Successful")				
			}
			if balanceError != nil {
				fmt.Println(balanceError.Error())
			}
			
		}
		
		if input == 3 {
			fmt.Println("\nYour Main balance is ", mainBalance)
			
		}
		
		if input == 4 {
			for i := range transHis {
				fmt.Println("\n ", transHis[i])
			}
		}
		
		if input == 5 {
			a += 2
		}
	}
}


func main() {
	fmt.Println("Hello, World!")
	
	
	atm(1000)
}

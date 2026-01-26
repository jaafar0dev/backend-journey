package main

import (
	"fmt"
	"errors"
	"time"
	"strconv"
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
				currentTime := time.Now()
				

				metaData := "You withdrew $" + strconv.Itoa(amount) + " at " + string(currentTime.String()) 
				transHis = append(transHis, metaData)
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
				
				currentTime := time.Now()
				
				metaData := "You made a transer of $" + strconv.Itoa(amount) + " at " + string(currentTime.String()) 
				transHis = append(transHis, metaData)				
			}
			if balanceError != nil {
				fmt.Println(balanceError.Error())
			}
			
		}
		
		if input == 3 {
			fmt.Println("\nYour Main balance is ", mainBalance)
			
			currentTime := time.Now()
			metaData := "You checked your balance at" + string(currentTime.String()) 
			transHis = append(transHis, metaData)
			
		}
		
		if input == 4 {
			for i := range transHis {
				fmt.Println(transHis[i])
			}
		}
		
		if input == 5 {
			a += 2
		}
	}
}

// FizzBuzz 

func fizzbuzz(){
	fmt.Print("Input a number:  ")
	
	var wrongInput error
	
	var number int
	fmt.Scan(&number)
	
	// if number != Type(int) {
		//wrongInput = errors.New("Only a number is allowed sorry!") }
		
	
	if wrongInput != nil{
		fmt.Print(wrongInput.Error())
	} else if number % 5 == 0 {
		fmt.Print("Fizz")
	} else if number % 3 == 0{
		fmt.Print("Buzz")
	} else if number % 3 == 0 && number % 5 == 0 {
		fmt.Print("FizzBuzz")
	} else {
		fmt.Print("Thanks")
	}
	
}

func adventureTime(lives int){
	a := 0
	var inp int
	
	
	for a < 1 {
		
		fmt.Println("Welcome Adventurer! My name is finn the human!")
		fmt.Println("Let's go on an adventure through this dungeon!")
		
		fmt.Println("What shall we do?\n")
		
		fmt.Println("Press 1 to enter the dungeon (Mathematical!")
		fmt.Println("Press 2 to escape! (Boooooo!)")
		
		fmt.Scan(&inp)
		
		if inp == 1{
			
			
			fmt.Println("Now we're in, there's a huge river, what should we do?")
			
			fmt.Println("Your options are:")
			fmt.Println("1. Stretch Jake across the river")
			fmt.Println("2. Swim scross the river yourself!")
			fmt.Println("3. Turn back")
			
			
		}
	
		if inp == 2{
			print("Let's leave Jake, this guy's a coward!")
			a += 1
		}
	}
}

func main() {
	adventureTime()
}


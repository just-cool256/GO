package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// In a real-world application, you would likely want to persist the account balance to a file or database.
// For simplicity, this example keeps it in memory only for the duration of the program run.
// Here's a simple function to write the balance to a file (not used in this example, but illustrative).

const balanceFileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err:= os.ReadFile(balanceFileName)
	if err != nil {
		return 1000.0, errors.New("Failed to find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.0, errors.New("Failed to parse balance from file.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("Starting with default balance of 1000.0")
		fmt.Println("------------------------------")
		// return
		// panic("Can't continue, sorry.")
		// this will stop the program immediately with message
	}

	// In go you only have one kind of loop - the for loop.
	// eg: for i:=0; i<10; i++ {
	//     //code to be executed
	// }
	// infinite loop
	// for {
	//     //code to be executed
	// }
	// Besides the for variations introduced in the last lectures, there also is another common variation (which will also be shown again later in the course):

	// for someCondition {
	//   //do something ...
	// }
	// // someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).

	// // In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// the switch case statement in Go works such that only one case can become active anyways. In other languages you need to add break at the end of each case to make sure that no other case gets triggered.
		// But still, the break keyword has this different meaning when used inside of a switch block.
		switch choice {
			case 1:
				fmt.Println("Your balance is ", accountBalance)
			case 2:
				var depositAmount float64
				fmt.Print("Enter amount to deposit: ")
				fmt.Scan(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Deposit amount must be positive!")
					continue
				}
				accountBalance += depositAmount
				fmt.Println("Deposit successful! New balance is ", accountBalance)
				writeBalanceToFile(accountBalance)
			case 3:
				var withdrawAmount float64
				fmt.Print("Enter amount to withdraw: ")
				fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("Withdraw amount must be positive!")
					continue
				}
				if withdrawAmount > accountBalance {
					fmt.Println("Insufficient funds!")
					continue
				}
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawal successful! New balance is ", accountBalance)
				writeBalanceToFile(accountBalance)
			case 4:
				fmt.Println("Thank you for using Go Bank. Goodbye!")
				return
				// break
			default:
				fmt.Println("Invalid choice!")
		}

		// Alternative without switch-case

		// if choice == 1 {
			// fmt.Println("Your balance is ", accountBalance)
		// } else if choice == 2 {
			// var depositAmount float64
			// fmt.Print("Enter amount to deposit: ")
			// fmt.Scan(&depositAmount)
			// if depositAmount <= 0 {
				// fmt.Println("Deposit amount must be positive!")
				// continue
			// }
			// accountBalance += depositAmount
			// fmt.Println("Deposit successful! New balance is ", accountBalance)
		// } else if choice == 3 {
			// var withdrawAmount float64
			// fmt.Print("Enter amount to withdraw: ")
			// fmt.Scan(&withdrawAmount)

			// if withdrawAmount <= 0 {
				// fmt.Println("Withdraw amount must be positive!")
				// continue
			// }
			// if withdrawAmount > accountBalance {
				// fmt.Println("Insufficient funds!")
			// } else {
				// accountBalance -= withdrawAmount
				// fmt.Println("Withdrawal successful! New balance is ", accountBalance)
			// }
		// } else if choice == 4 {
			// fmt.Println("Thank you for using Go Bank. Goodbye!")
			// break
		// } else {
			// fmt.Println("Invalid choice!")
		// }
	}

}

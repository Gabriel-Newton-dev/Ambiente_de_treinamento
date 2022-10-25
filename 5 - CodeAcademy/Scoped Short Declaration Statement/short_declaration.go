// 1- In main.go we have an if statement and a switch statement that can each use the short declaration.
// Let’s first change the if statement to declare success inside the statement. You’ll have to delete the provided success variable.

package main

import "fmt"

func main() {

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000
	numOfThieves := 5

	switch numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}

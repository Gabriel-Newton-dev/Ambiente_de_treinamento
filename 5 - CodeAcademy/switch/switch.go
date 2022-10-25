// 1.Add a switch statement that takes the condition name.
// 2.Add one case statement with the value "Butch" and inside its block print the string "Head to Robbers Roost.".
// 3.Add a second case statement with the value "Bonnie" and inside its block, print the string "Stay put in Joplin.".
// 4.Add a default statement that will print out "Just hide!".

package main

import "fmt"

func main() {
	name := "H. J. Simp"

	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}

	// Add your switch statement below:

}

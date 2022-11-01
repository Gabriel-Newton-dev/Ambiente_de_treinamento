package main

import "fmt"

var name string

func main() {

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	Welcome(name)

}

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome tom my party %s", name)
}

// HappyBirthday wishes happy birthday to the birtday person and inform your age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	info := Welcome(name) + "\n"
	info += fmt.Sprintf("You have been assigned to table %03d. ", table)
	info += fmt.Sprintf("Your table is %s, exactly %.1f meters from here.\n", direction, distance)
	info += fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return info
}

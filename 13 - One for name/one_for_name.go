package main

import "fmt"

func main() {
	name := "Gabriel"
	ShareWith(name)
}

// ShareWith needs a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

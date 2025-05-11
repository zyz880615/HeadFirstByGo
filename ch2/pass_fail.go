// pass_fail reports whether a grade is passing or failing.
package main

import (
	"HeadFirstByGo/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a grade: ")

	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}

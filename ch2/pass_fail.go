// pass_fail reports whether a grade is passing or failing.
package main

import (
	"HeadFirstByGo/ch4/keyboard"
	"fmt"
	//"github.com/zhangyunzheng/HeadFirstByGo/ch4/keyboard"
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

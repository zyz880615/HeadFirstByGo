// average calculates the average of several numbers.
package main

import (
	"fmt"
	"github.com/zyz880615/HeadFirstByGo/ch5/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("/Users/zhangyunzheng/Documents/GoProject/HeadFirstByGo/ch5/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average %0.2f\n", sum/sampleCount)

}

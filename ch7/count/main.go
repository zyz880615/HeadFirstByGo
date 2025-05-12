// count tallies the number of times each line
// occurs within a file.
package main

import (
	"fmt"
	"github.com/zyz880615/HeadFirstByGo/ch5/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("/Users/zhangyunzheng/Documents/GoProject/HeadFirstByGo/ch7/vote.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}

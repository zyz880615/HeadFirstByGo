// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")
	//读取键盘输入
	reader := bufio.NewReader(os.Stdin)
	//读取用户输入内容，直到按了<Enter>键
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//从input字符串中删除换行字符
	input = strings.TrimSpace(input)
	//将字符串转换为float64值
	grade, err := strconv.ParseFloat(input, 64)

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

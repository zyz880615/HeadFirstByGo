package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//获取当前日期和时间的整数形式
	seconds := time.Now().Unix()
	//随机数生成器
	rand.NewSource(seconds)
	//生成1-100的随机数
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	
	//从键盘读取输入
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Make a guess: ")
	//读取用户输入的内容，直到按下<Enter>
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//删除换行符
	input = strings.TrimSpace(input)
	//将输入字符串转换为整数
	guess, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
	}

}

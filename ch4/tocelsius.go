// tocelsius converts a temperature from Fahrenheit to Celsius.
package main

import (
	"HeadFirstByGo/ch4/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	//获取一个温度
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	//将温度转换为摄氏度
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degress Celsius\n", celsius)
}

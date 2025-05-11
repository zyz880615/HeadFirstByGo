// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	//打开所提供的文件名
	file, err := os.Open(fileName)

	if err != nil {
		return numbers, err
	}

	//这个变量将跟踪我们应该赋值给哪个数组索引
	i := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//将文件行字符串转换为float64
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()

	//如果关闭文件时出现错误，返回该错误
	if err != nil {
		return numbers, err
	}

	//如果扫描文件出现错误，返回该错误
	if scanner.Err() != nil {
		return numbers, err
	}

	return numbers, nil

}

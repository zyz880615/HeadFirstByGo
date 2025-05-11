// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	//打开所提供的文件名
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//将文件行字符串转换为float64
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		//追加新的数字给切片
		numbers = append(numbers, number)
	}

	err = file.Close()

	//如果关闭文件时出现错误，返回该错误
	if err != nil {
		return nil, err
	}

	//如果扫描文件出现错误，返回该错误
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil

}

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) {
	//打印当前目录
	fmt.Println(path)

	//获取包含目录内容的切片
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		//用斜杠将目录路径和文件名连接起来
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() { //如果这是一个子目录
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}

}

func main() {
	defer reportPanic()
	scanDirectory("/Users/zhangyunzheng/Documents/GoProject/HeadFirstByGo/ch12")
}

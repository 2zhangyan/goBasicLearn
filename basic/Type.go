package main

import "fmt"

//类型转换
func conversion() {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为：%f \n", mean)
}

func main() {
	conversion()
}

package typefunc

import "fmt"

// 声明了一个函数类型
type testInt func(int) bool

// isOdd 是否奇数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

// isEven 是否为偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// TypeFunc 类型函数
func TypeFunc(slice []int) {

	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd) // 函数当作值来传递了

	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven) // 函数当作值来传递了

	fmt.Println("Even elements of slice are: ", even)
}

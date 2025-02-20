package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stack []string // 使用切片来模拟栈
	reader := bufio.NewReader(os.Stdin)

	// 读取 n
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// 读取 n 个数字
	line2, _ := reader.ReadString('\n')
	values := strings.Fields(line2)

	for i := 0; i < n; i++ {
		stack = append(stack, values[i]) // 压入栈

		// 当栈顶 4 个元素相同时，移除它们
		for len(stack) >= 3 && allEqual(stack[len(stack)-3:]) {
			stack = stack[:len(stack)-3]
		}
	}

	// 输出结果
	if len(stack) == 0 {
		fmt.Println("0")
	} else {
		fmt.Println(strings.Join(stack, " "))
	}
}

// allEqual 检查切片中的所有元素是否相等
func allEqual(slice []string) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[0] {
			return false
		}
	}
	return true
}

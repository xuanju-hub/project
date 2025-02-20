package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// 创建一个标准输入读取器
	reader := bufio.NewReader(os.Stdin)

	// 读取第一行输入
	firstLine, _ := reader.ReadString('\n')
	// 去掉末尾的换行符
	firstLine = strings.TrimSpace(firstLine)

	// 解析第一行的两个整数
	firstParts := strings.Split(firstLine, " ")
	n, _ := strconv.Atoi(firstParts[0])
	m, _ := strconv.Atoi(firstParts[1])

	// 读取第二行输入
	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)

	// 解析第二行的整数数组
	numStrings := strings.Split(secondLine, " ")
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(numStrings[i])
	}

	check := func(x int) bool {
		sum := 1
		cap := x
		for i := 0; i < len(numbers); {
			if cap-numbers[i] >= 0 {
				cap -= numbers[i]
				i++
			} else {
				sum++
				cap = x
			}
		}

		return sum <= m
	}

	left, right := slices.Max(numbers)-1, 50000

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Println(right)
}

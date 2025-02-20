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

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	if sum < m {
		fmt.Println("-1528")
		return
	}

	check := func(x int) bool {
		sum := 0
		for _, v := range numbers {
			sum += min(x, v)
		}

		return sum <= m
	}

	left, right := 0, slices.Max(numbers)+1

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Println(left)

}

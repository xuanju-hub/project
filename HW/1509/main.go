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
	firstLine = strings.TrimSpace(firstLine)

	numStrings := strings.Split(firstLine, " ")
	n := len(numStrings)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(numStrings[i])
	}

	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimSpace(secondLine)
	m, _ := strconv.Atoi(secondLine)

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
	if left == slices.Max(numbers) {
		fmt.Println("-1528")
	} else {
		fmt.Println(left)
	}
}

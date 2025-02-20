package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)

	numStrings := strings.Split(line1, " ")
	n := len(numStrings)
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	m, _ := strconv.Atoi(line2)
	// 初始化二分查找的边界
	l, r := 1, n
	p := (l + r) / 2
	path := "S" // 记录搜索路径

	target := m
	// 二分查找
	for p != target {
		if target < p {
			r = p - 1
			path += "L"
		} else {
			l = p + 1
			path += "R"
		}
		p = (l + r) / 2
	}

	// 目标找到，添加 "Y"
	path += "Y"
	fmt.Println(path)
}

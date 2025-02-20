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

	// 读取第一行输入
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	numStrings := strings.Split(firstLine, ",")
	n := len(numStrings)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(numStrings[i])
	}

	m := make(map[int]int)
	for _, v := range numbers {
		m[v]++
	}

	for k, v := range m {
		if v > n/2 {
			fmt.Println(k)
			return
		}
	}

	fmt.Println("0")
}

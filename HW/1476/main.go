package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	str1 := strings.Split(strings.TrimSpace(line1), " ")

	stack := []string{}
	res := []string{}
	for _, v := range str1 {
		if x := slices.Index(stack, v); x >= 0 {
			for i := len(stack) - 1; i >= x; i-- {
				res = append(res, stack[i])
			}
			stack = stack[:x]
		}
		stack = append(stack, v)
	}

	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	for i, v := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

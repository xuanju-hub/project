package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	str1 := strings.Split(strings.TrimSpace(line1), "/")

	line2, _ := reader.ReadString('\n')
	str2 := strings.Split(strings.TrimSpace(line2), " ")

	str3 := strings.Split(str2[1], "/")

	maxn := 0
	for _, v := range str3 {
		if v == "." || v == "" {
			continue
		}

		if v == ".." {
			str1 = str1[:len(str1)-1]

		} else {
			str1 = append(str1, v)
			if len(str1)-1 > maxn {
				maxn = len(str1) - 1
			}
		}
	}

	for i, v := range str1 {
		if i > 0 {
			fmt.Print("/")
		}
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(maxn)

}

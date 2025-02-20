package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)
	arrStr := strings.Split(input, ",")
	arr := make([]int, len(arrStr))
	for i, numStr := range arrStr {
		arr[i], _ = strconv.Atoi(numStr)
	} //处理输入

	sort.Ints(arr) //排序
	i := 0
	ans1 := -1
	ans2 := -1
	for i < len(arr) { //第一层循环
		cnt := 1
		for j := i + 1; j < len(arr); j++ { //记录最大相邻数
			if arr[j]-arr[j-1] == 1 { //向后找有没有相邻的
				cnt++
			} else {
				break
			}
		}
		if cnt > ans2 { //更新答案
			ans2 = cnt
			ans1 = arr[i]
		}
		i += cnt //按照题解意思,我们后面需要查找的下标为i+cnt
	}
	fmt.Printf("%d,%d", ans1, ans2)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// 读取 s 和 y
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	s, _ := strconv.Atoi(line1) // 图片总数
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	y, _ := strconv.Atoi(line2) // 经验值排名

	// 记录出现过的有效分数（去重）
	scoreSet := make(map[int]bool)
	// 存储所有有效分数（用于排序）
	var scores []int
	// 记录分数 -> 图片编号（map 结构）
	scoreMap := make(map[int][]int)

	// 读取 s 个图片评分
	line, _ := reader.ReadString('\n')
	values := strings.Fields(line)

	for i := 1; i <= s; i++ {
		x, _ := strconv.Atoi(values[i-1])

		if x <= 20 {
			continue
		}

		if scoreSet[x] {
			scoreMap[x] = append(scoreMap[x], i)
		} else {
			scoreSet[x] = true
			scores = append(scores, x)
			scoreMap[x] = []int{i}
		}
	}

	// 对所有有效分数按 **降序** 排序
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	// **获取排名第 y 的分数**

	ans := min(len(scores), y)

	targetScore := scores[ans-1]
	fmt.Print(targetScore, " ")

	// **输出该分数对应的所有图片编号**
	for i, v := range scoreMap[targetScore] {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println() // 换行
}

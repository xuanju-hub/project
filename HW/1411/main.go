package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 定义套餐结构体
type Package struct {
	mainID int // 主食 ID
	cost   int // 成本
	profit int // 利润
	index  int // 原始索引
}

// 自定义排序规则（成本优先，其次利润，最后索引）
func sortPackages(packages []Package) {
	sort.Slice(packages, func(i, j int) bool {
		if packages[i].cost != packages[j].cost {
			return packages[i].cost < packages[j].cost // 成本小的优先
		}
		if packages[i].profit != packages[j].profit {
			return packages[i].profit > packages[j].profit // 利润高的优先
		}
		return packages[i].index < packages[j].index // 索引小的优先
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 读取 M（套餐数量）
	Mstr, _ := reader.ReadString('\n')
	M, _ := strconv.Atoi(strings.TrimSpace(Mstr))

	// 读取 M 组套餐
	packages := make([]Package, M)
	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Fields(line)
		packages[i].mainID, _ = strconv.Atoi(values[0])
		packages[i].cost, _ = strconv.Atoi(values[1])
		packages[i].profit, _ = strconv.Atoi(values[2])
		packages[i].index = i
	}

	// 读取 N（主食限制数量）
	Nstr, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(Nstr))

	// 读取 N 组主食限制
	limitMap := make(map[int]int) // 记录主食 ID -> 限制数量
	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Fields(line)
		id, _ := strconv.Atoi(values[0])
		limit, _ := strconv.Atoi(values[1])
		limitMap[id] = limit
	}

	// 读取 NUM（最多选择的套餐数量）
	NUMstr, _ := reader.ReadString('\n')
	NUM, _ := strconv.Atoi(strings.TrimSpace(NUMstr))

	// **按排序规则对套餐排序**
	sortPackages(packages)

	// **选择套餐**
	selected := []int{}              // 存储最终选择的套餐索引
	selectedMap := make(map[int]int) // 记录每种主食已选择的数量

	for _, pkg := range packages {
		// 如果该主食有选择限制
		if limit, exists := limitMap[pkg.mainID]; exists {
			if selectedMap[pkg.mainID] < limit {
				selected = append(selected, pkg.index)
				selectedMap[pkg.mainID]++
			}
		} else {
			// 没有限制的套餐直接选择
			selected = append(selected, pkg.index)
		}

		// 选满 NUM 组套餐
		if len(selected) == NUM {
			break
		}
	}

	// **输出结果**
	if len(selected) == NUM {
		sort.Ints(selected) // 按原始输入顺序排序
		for i, idx := range selected {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(idx)
		}
		fmt.Println()
	} else {
		fmt.Println("-1")
	}
}

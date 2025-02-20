package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Package struct {
	time   string
	client string
	factor string
	num    int
}

func sortPackages(packages []Package) {
	sort.Slice(packages, func(i, j int) bool {
		if packages[i].time != packages[j].time {
			return packages[i].time < packages[j].time
		}
		if packages[i].client != packages[j].client {
			return packages[i].client < packages[j].client
		}
		return packages[i].factor < packages[j].factor
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 读取 M（套餐数量）
	nstr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nstr))

	// 读取 M 组套餐
	packages := make([]Package, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Split(strings.TrimSpace(line), ",")
		packages[i].time = values[0]
		packages[i].client = values[1]
		packages[i].factor = values[2]
		x, _ := strconv.Atoi(values[3])
		if x < 0 {
			x = 0
		}
		packages[i].num = x
	}
	mstr, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mstr))

	factorMap := make(map[string]int)

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(values[1])
		if x < 0 {
			x = 0
		}
		factorMap[values[0]] = x
	}

	sortPackages(packages)

	clientSet := make(map[string]int)

	for i := 0; i < n; i++ {

		clientSet[packages[i].client] += factorMap[packages[i].factor] * packages[i].num
	}

	keys := make([]string, 0, len(clientSet))
	for k := range clientSet {
		keys = append(keys, k)
	}

	sort.Strings(keys) // 按字母顺序排序

	for _, k := range keys {
		fmt.Printf("%s,%d", k, clientSet[k])
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"slices"
)

// 求交集
// 定义一个函数intersection，用于找出两个整数切片的交集
func intersection(slice1, slice2 []int) []int {
    // 初始化一个空的整数切片result，用于存储交集结果
	result := []int{}
    // 初始化一个map，用于存储slice1中的元素，map的键为整数，值为布尔类型
	set := make(map[int]bool)

    // 遍历slice1中的每个元素
	for _, v := range slice1 {
        // 将slice1中的元素作为键存入map中，值为true，表示该元素存在于slice1中
		set[v] = true
	}

    // 遍历slice2中的每个元素
	for _, v := range slice2 {
        // 检查当前元素是否存在于map中，即是否也在slice1中
		if set[v] {
            // 如果存在，则将该元素添加到result切片中
			result = append(result, v)
		}
	}

    // 返回结果切片result，即两个切片的交集
	return result
}

// 求并集
func union(slice1, slice2 []int) []int {
	result := []int{}
	set := make(map[int]bool)

	for _, v := range slice1 {
		if !set[v] {
			result = append(result, v)
			set[v] = true
		}
	}

	for _, v := range slice2 {
		if !set[v] {
			result = append(result, v)
			set[v] = true
		}
	}

	return result
}

// 求差集 (slice1 - slice2)
// difference 函数用于计算两个整数切片的差集
// 参数 slice1 和 slice2 分别表示两个整数切片
// 返回值是一个新的切片，包含在 slice1 中但不在 slice2 中的元素
func difference(slice1, slice2 []int) []int {
	// 初始化结果切片，用于存储差集元素
	result := []int{}
	// 使用 map 来记录 slice2 中的元素，map 的键为元素值，值为布尔类型
	set := make(map[int]bool)

	// 遍历 slice2，将每个元素作为键存入 map 中，值为 true
	for _, v := range slice2 {
		set[v] = true
	}

	// 遍历 slice1，检查每个元素是否在 map 中
	for _, v := range slice1 {
		// 如果元素不在 map 中，说明该元素不在 slice2 中
		if !set[v] {
			// 将该元素添加到结果切片中
			result = append(result, v)
		}
	}
	// 返回结果切片，即差集
	return result
}

func main() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{4, 5, 6, 7, 8}
	fmt.Println("交集:", intersection(sliceA, sliceB))
	fmt.Println("并集:", union(sliceA, sliceB))
	fmt.Println("差集 (A - B):", difference(sliceA, sliceB))
	fmt.Println("差集 (B - A):", difference(sliceB, sliceA))
	slices.Sort(sliceA)
	// 交集: [4 5]
	// 并集: [1 2 3 4 5 6 7 8]
	// 差集 (A - B): [1 2 3]
	// 差集 (B - A): [6 7 8]
}

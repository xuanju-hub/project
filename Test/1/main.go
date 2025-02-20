package main

import (
	"fmt"
)

// 求交集
func intersection(slice1, slice2 []int) []int {
	result := []int{}
	set := make(map[int]bool)

	for _, v := range slice1 {
		set[v] = true
	}

	for _, v := range slice2 {
		if set[v] {
			result = append(result, v)
		}
	}

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
func difference(slice1, slice2 []int) []int {
	result := []int{}
	set := make(map[int]bool)

	for _, v := range slice2 {
		set[v] = true
	}

	for _, v := range slice1 {
		if !set[v] {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{4, 5, 6, 7, 8}
	fmt.Println("交集:", intersection(sliceA, sliceB))
	fmt.Println("并集:", union(sliceA, sliceB))
	fmt.Println("差集 (A - B):", difference(sliceA, sliceB))
	fmt.Println("差集 (B - A):", difference(sliceB, sliceA))
	// 交集: [4 5]
	// 并集: [1 2 3 4 5 6 7 8]
	// 差集 (A - B): [1 2 3]
	// 差集 (B - A): [6 7 8]
}

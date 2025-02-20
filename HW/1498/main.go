package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxn = 1010

func main() {
	var stack [maxn]int // 栈数组
	top := 0            // 栈顶指针

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	values := strings.Fields(line)

	for i := 0; i < len(values); i++ {
		x, _ := strconv.Atoi(values[i])
		for merged := true; merged; {
			merged = false // 每轮开始时重置合并标志
			sum := 0       // 计算连续元素之和

			// 从栈顶向下遍历，查找是否有一段和等于 x
			for i := top; i > 0; i-- {
				sum += stack[i]

				if sum == x { // 找到一段和等于 x 的元素
					x += sum      // 更新 x（相当于合并）
					top = i - 1   // 栈顶调整，移除合并部分
					merged = true // 记录发生了合并
					break         // 重新检查新的 x 是否还能继续合并
				}
			}
		}
		top++
		stack[top] = x
	}
	//for {
	//	_, err := fmt.Scanf("%d", &x) // 读入数据
	//	if err != nil {
	//		break
	//	}
	//	for {
	//		flag := false              // 标记是否存在合并的情况
	//		var tmp int                // 临时变量存储和
	//		for i := top; i > 0; i-- { // 暴力查找是否和为x的数
	//			tmp += stack[i]
	//			if tmp == x { // 存在一段和为x，合并，并更新栈
	//				x = x + tmp
	//				top = i - 1
	//				flag = true
	//				break
	//			}
	//		}
	//		if !flag {
	//			break
	//		}
	//	}
	//	top++
	//	stack[top] = x
	//}

	for top > 0 {
		fmt.Printf("%d ", stack[top])
		top--
	}
}

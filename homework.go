package main

import (
	"fmt"
	"sync"
)

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func NumAutoAdd(num *int) {
	*num += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func ItemIncrement(num *[]int) {
	for i := 0; i < len(*num); i++ {
		(*num)[i] = (*num)[i] * 2
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func WorkFormRouting() {
	fmt.Println("=== go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数 ===")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Printf("偶数: ")
		for i := 2; i < 10; i += 2 {
			fmt.Printf(" %d", i)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Printf("\n奇数: ")
		for i := 1; i < 10; i += 2 {
			fmt.Printf(" %d", i)
		}
	}()
	wg.Wait()
	fmt.Println("\n所有数字输出完成")
}

func main() {
	//fmt.Println("开始测试指针加10")
	//x := 10
	//fmt.Printf("指针运算：原值 %d\n", x)
	//NumAutoAdd(&x)
	//fmt.Printf("指针运算：加10 %d\n", x)
	//fmt.Println("开始测试数组x2")
	//s := []int{1, 2, 3}
	//fmt.Printf("原始切片：%v \n", s)
	//ItemIncrement(&s)
	//fmt.Printf("切片元素x2：%v\n", s)
	WorkFormRouting()
}

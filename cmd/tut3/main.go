package main

import (
	"fmt"
)

func Swap(num []int, idx int) {
	num[idx], num[idx+1] = num[idx+1], num[idx]
}

func BuubbleSort(num []int) {
	n := len(num)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if num[j] > num[j+1] {
				Swap(num, j)
			}
		}
	}

}

func main() {
	//bubble sort implementation
	fmt.Println("Enter 10 integers into the array:")
	num := []int{}
	for i := 0; i < 10; i++ {
		x := 0
		fmt.Scanf("%d\n", &x)
		num = append(num, x)
	}
	BuubbleSort(num)
	fmt.Println("Sorted Array:\t", num)
}

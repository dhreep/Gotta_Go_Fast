package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ReadValues() (numbers []int, err error) {
	fmt.Println("Please input numbers(separate with space): ")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine()

	ns := strings.Split(string(a), " ")
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return
}

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}

func Sort(numbers []int, wg *sync.WaitGroup) {
	QuickSort(numbers)

	if wg != nil {
		wg.Done()
	}
}

func CreatePartitions(numbers []int, k int) (partitions [][]int) {
	var n = len(numbers)
	partitions = make([][]int, k)

	lenPartition := n / k

	for j := range partitions {
		partitions[j] = numbers[j*lenPartition : (j+1)*lenPartition]
	}

	return
}

func MergePartitions(s1 []int, s2 []int) (merged []int) {
	var n1, n2 int = len(s1), len(s2)

	if n1 < n2 {
		s1, s2 = s2, s1
		n1, n2 = n2, n1
	}

	merged = make([]int, n1+n2)

	i, j, k := 0, 0, 0
	for {
		if s1[i] < s2[j] {
			merged[k] = s1[i]
			i++
			k++
		} else if s1[i] == s2[j] {
			merged[k] = s1[i]
			i++
			k++

			merged[k] = s2[j]
			j++
			k++
		} else if s1[i] > s2[j] {
			merged[k] = s2[j]
			j++
			k++
		}

		if j == n2 {
			copy(merged[k:], s1[i:])
			k = n1 + n2
		}

		if i == n1 {
			copy(merged[k:], s2[j:])
			k = n1 + n2
		}

		if k == (n1 + n2) {
			break
		}
		fmt.Println(merged)
	}

	return
}

func Merge(partitions [][]int, k int) (sorted []int) {

	var psorted [][]int

	for i := 0; i <= k/2; i += 2 {
		psorted = append(psorted, MergePartitions(partitions[i], partitions[i+1]))
	}

	if len(psorted) == 2 {
		sorted = MergePartitions(psorted[0], psorted[1])
	} else {
		Merge(psorted, k)
	}
	return
}

func main() {
	var k = 4
	var wg sync.WaitGroup

	if numbers, err := ReadValues(); err != nil {
		log.Fatal("Error reading text from console: ", err)
	} else {
		nTotal := len(numbers)

		if nTotal <= k {
			QuickSort(numbers)
			fmt.Printf("Sorted numbers: %v", numbers)
		} else {

			partitions := CreatePartitions(numbers, k)

			wg.Add(k)
			for i := 0; i < k; i++ {
				go Sort(partitions[i], &wg)
			}
			wg.Wait()

			sorted := Merge(partitions, k)

			fmt.Printf("Sorted numbers: %v", sorted)
		}
	}
}

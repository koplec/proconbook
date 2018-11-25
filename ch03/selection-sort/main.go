package main

import "fmt"

func selectionSort(a []int) (switchCount int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		minj := i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		if i != minj {
			a[i], a[minj] = a[minj], a[i]
			switchCount++
		}
	}
	return switchCount
}

func main() {
	var n int
	var a []int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	a = make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &a[i])
		if err != nil {
			panic(err)
		}
	}
	sc := selectionSort(a)
	trace(a)
	fmt.Printf("%d\n", sc)
}

/**
 * 配列の要素を順番に出力
 */
func trace(a []int) {
	for i := 0; i < len(a); i++ {
		if i > 0 { //隣接する要素の間に１つの空白を出力
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
}

package main

import "fmt"

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

/**
 *挿入ソート
 */
func insertionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		trace(a)
	}
	return
}

func main() {
	var n int
	var a []int

	_, err := fmt.Scanf("%d", &n) //scan array size
	if err != nil {
		panic(err)
	}
	a = make([]int, n)
	for i := 0; i < n; i++ { //scan array like "5 2 4 6 1 3"
		_, err = fmt.Scanf("%d", &a[i])
		if err != nil {
			panic(err)
		}
	}

	trace(a)
	insertionSort(a)

	return
}

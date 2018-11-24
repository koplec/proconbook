package main

import "fmt"

func main() {
	var n int
	var a []int
	_, err := fmt.Scanf("%d", &n)
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

	sc := bubbleSort(a)
	trace(a)
	fmt.Printf("%d\n", sc)
	return
}

func bubbleSort(a []int) (switchCount int) {
	flag := true //contains reversed elements
	n := len(a)
	for i := 0; flag; i++ {
		flag = false
		for j := n - 1; j >= i+1; j-- {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
				switchCount++
			}
		}
	}
	return switchCount
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

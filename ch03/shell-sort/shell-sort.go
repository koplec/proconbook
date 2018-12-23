package main

import "fmt"
import "sort"

func main() {
	a := []int{4, 8, 9, 1, 10, 6, 2, 5, 3, 7}
	fmt.Printf("before sort:%v\n", a)
	count := shellSort(a, len(a))
	fmt.Printf("after sort:%v\n", a)
	fmt.Printf(" count:%d\n", count)
}

/**
 * n:配列の要素数
 * g: いくつずつとばしてinsertion-sortを実施するか？
 * g = 1なら通常のinsertion-sort
 * 実際に挿入を実施した回数を返す
 */
func insertionSort(A []int, n, g int) (count int) {
	count = 0
	for i := g; i < n; i++ { //iは増えていく
		v := A[i]
		//jはiより小さいインデックス
		j := i - g
		//v=A[i]より大きな要素を後ろに持っていく
		for j >= 0 && A[j] > v { //jは減っていく
			//最初のループで、
			//  j+g=(i-g)+g=iになること注意
			//  j = i - g
			//前の値で後ろの値を更新する
			//A[j] <= vが見つかった時点でループは終了
			A[j+g] = A[j]
			j = j - g
			count++
		}
		//j+g番目はiから負の方向にgずつ進んでいったときに0より大きい一番最小の数
		//g=1のときは、j+g=0
		A[j+g] = v
	}
	return count
}

func shellSort(A []int, n int) (count int) {
	count = 0
	gs := make([]int, 0)
	//g_i+1 = 3*g_i + 1の数列が効率がいいことが分かっている( O n^1.25)程度
	for g := 1; g < n; g = 3*g + 1 {
		gs = append(gs, g)
	}
	sort.Reverse(sort.IntSlice(gs))

	fmt.Printf("gs=%v\n", gs)
	for i := 0; i < len(gs); i++ {
		count += insertionSort(A, n, gs[i])
	}

	return
}

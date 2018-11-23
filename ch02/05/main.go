package main

import "math"
import "fmt"

const max_index int = 200000

func main() {
	var r = make([]int, max_index)
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)

	}
	fmt.Printf("n = %d\n", n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&r[i])
		if err != nil {
			panic(err)
		}
	}

    maxv := math.MinInt64
    minv := r[0]

    for i := 1; i < n; i++{
        maxv = max(maxv, r[i] - minv)
        minv = min(minv, r[i])
    }

    fmt.Printf("max value=%d\n", maxv)
    
    return
}

func max(x int, y int) int{
    if x < y {
        return y
    }else{
        return x
    }
}

func min(x int, y int) int{
    if x > y {
        return y
    }else{
        return x
    }
}

package main

import "errors"
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

const QUEUE_MAX = 1000

type IntQueue struct {
	Top int
	S   [QUEUE_MAX + 1]int //0番目には何も値を入れない
}

func NewIntQueue() *IntQueue {
	return &IntQueue{
		Top: 0,
	}
}

func (q *IntQueue) IsEmpty() bool {
	return q.Top == 0
}
func (q *IntQueue) IsFull() bool {
	return q.Top == QUEUE_MAX
}
func (q *IntQueue) Push(x int) (err error) {
	if q.IsFull() {
		err = errors.New("overflow")
		return err
	} else {
		q.Top++
		q.S[q.Top] = x
		return nil
	}
}

func (q *IntQueue) Pop() (int, error) {
	if q.IsEmpty() {
		err := errors.New("underflow")
		return 0, err
	} else {
		q.Top--
		return q.S[q.Top+1], nil
	}
}

func main() {
	line, _ := readLine()
	q := NewIntQueue()

	fmt.Printf("scan line -> %s\n", line)
	strs := strings.Split(line, " ")
	for idx, s := range strs {
		fmt.Printf("%d -> %s\n", idx, s)
		switch s {
		case "+":
			a, _ := q.Pop()
			b, _ := q.Pop()
			q.Push(a + b)
		case "*":
			a, _ := q.Pop()
			b, _ := q.Pop()
			q.Push(a * b)
		case "-":
			a, _ := q.Pop()
			b, _ := q.Pop()
			q.Push(b - a)
		default:
			a, _ := strconv.Atoi(s)
			q.Push(a)
		}
	}

	ret, _ := q.Pop()
	fmt.Printf("%d\n", ret)
}

var rdr = bufio.NewReaderSize(os.Stdin, 100000)

/**
 * 長い文字列はfmt.Scanfだと時間がかかる
 */
func readLine() (string, error) {
	buf := make([]byte, 0, 100000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			return "", e
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf), nil
}

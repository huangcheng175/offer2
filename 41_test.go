package offer2

import (
	"fmt"
	"testing"
)

type MovingAverage struct {
	queue []int
	size  int
	sum   int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		queue: []int{},
		size:  size,
		sum:   0,
	}
}

func (m *MovingAverage) next(val int) float64 {
	m.queue = append(m.queue, val)
	m.sum += val
	if len(m.queue) <= m.size {
		return float64(m.sum) / float64(len(m.queue))
	}

	m.sum -= m.queue[0]
	m.queue = m.queue[1:]

	return float64(m.sum) / float64(len(m.queue))
}

func TestMovingAverage(t *testing.T) {
	v := Constructor(3)
	fmt.Println(v.next(1))
	fmt.Println(v.next(2))
	fmt.Println(v.next(3))
	fmt.Println(v.next(4))
	fmt.Println(v.next(10))

}

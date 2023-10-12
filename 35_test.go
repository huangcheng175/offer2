package offer2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {

	idx := make([]bool, 24*60)

	var (
		id  int
		err error
	)
	for _, t := range timePoints {
		id, err = calTimeIndex(t)
		if err != nil {
			return -1
		}
		if idx[id] {
			return 0
		}
		idx[id] = true
	}

	var p1 int
	for i := 0; i < len(timePoints); i++ {
		if idx[i] {
			p1 = i
			break
		}
	}

	minInterval := math.MaxInt
	start := p1
	p2 := p1 + 1
	for p2 < 24*60 {
		if !idx[p2] {
			p2++
			continue
		}
		minInterval = min(minInterval, p2-p1)
		p1 = p2
		p2 = p2 + 1
	}

	minInterval = min(minInterval, start+1440-p1)

	return minInterval

}

func min(a, b int) int {
	if a < b {
		return
	}
}
func calTimeIndex(t string) (id int, err error) {
	v := strings.Split(t, ":")
	if len(v) != 2 {
		err = errors.New("invalid time format")
		return
	}

	hour, err := strconv.ParseInt(v[0], 10, 64)
	if err != nil {
		return
	}
	minu, err := strconv.ParseInt(v[1], 10, 64)
	if err != nil {
		return
	}

	id = int(hour*60 + minu)

	return

}

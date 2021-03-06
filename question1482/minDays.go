package question1482

import (
	"math"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if k*m > n {
		return -1
	}
	minDay, maxDay := math.MaxInt32, 0
	for _, day := range bloomDay {
		if day < minDay {
			minDay = day
		}
		if day > maxDay {
			maxDay = day
		}
	}
	return minDay + sort.Search(maxDay-minDay, func(days int) bool {
		days += minDay
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}

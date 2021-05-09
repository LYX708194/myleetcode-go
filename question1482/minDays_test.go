package question1482

import "testing"

type question struct {
	bloomDay []int
	m        int
	k        int
}

func TestMinDays(t *testing.T) {
	questions := []question{
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        1,
		},
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        2,
		},
		{
			bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
			m:        2,
			k:        3,
		},
		{
			bloomDay: []int{1000000000, 1000000000},
			m:        1,
			k:        1,
		},
		{
			bloomDay: []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
			m:        4,
			k:        2,
		},
	}
	ans := []int{3, -1, 12, 1000000000, 9}
	for i := 0; i < len(questions); i++ {
		if minDays(questions[i].bloomDay, questions[i].m, questions[i].k) != ans[i] {
			t.Error()
		}
	}
}

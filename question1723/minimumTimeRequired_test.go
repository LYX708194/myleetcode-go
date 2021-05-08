package question1723

import (
	"testing"
)

type question struct {
	jobs []int
	k    int
}

func TestMinimumTimeRequired(t *testing.T) {
	questions := []question{
		{
			jobs: []int{3, 2, 3},
			k:    3,
		},
		{
			jobs: []int{1, 2, 4, 7, 8},
			k:    2,
		},
	}
	ans := []int{3, 11}
	for i := 0; i < len(questions); i++ {
		if minimumTimeRequired(questions[i].jobs, questions[i].k) != ans[i] {
			t.Error()
		}
	}
}

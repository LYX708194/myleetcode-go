package question1486

import (
	"testing"
)

type question struct {
	n int
	start int
}

func TestXorOperation(t *testing.T){
	questions := []question{
		{5,0},
		{4,3},
		{1,7},
		{10,5},
	}
	ans := []int{8,8,7,2}
	for i := 0;i < len(questions) ; i++ {
		if xorOperation(questions[i].n,questions[i].start) != ans[i]{
			t.Fail()
		}
	}
}

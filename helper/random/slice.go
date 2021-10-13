package randomHelper

import (
	"math/rand"
	"time"
)

func Numbers(start, end, count int) (nums []int) {
	nums = make([]int, 0)
	if end < start || (end-start) < count {
		return
	}
	var (
		num    int
		exists bool
	)
	if len(nums) < count {
		num = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(end-start) + start
		exists = false
		for _, v := range nums {
			if v == num {
				exists = true
				break
			}
		}
		if !exists {
			nums = append(nums, num)
		}
	}
	return
}

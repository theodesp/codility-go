package stacks_queues

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestFish(c *C) {
	c.Assert(Fish([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}), Equals, 2)
}

func (s *MySuite) BenchmarkFish(c *C) {
	for i := 0; i < c.N; i++ {
		Fish([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0})
	}
}
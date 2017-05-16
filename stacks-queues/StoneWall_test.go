package stacks_queues

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestStonewall(c *C) {
	c.Assert(Stonewall([]int{8, 8, 5, 7, 9, 8, 7, 4, 8}) ,Equals, 7)
}

func (s *MySuite) BenchmarkStonewall(c *C) {
	for i := 0; i < c.N; i++ {
		Stonewall([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})
	}
}
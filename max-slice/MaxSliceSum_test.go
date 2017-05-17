package max_slice

import (
	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxSliceSum(c *C) {
	c.Assert(MaxSliceSum([]int{-10}), Equals, -10)
}

func (s *MySuite) BenchmarkMaxSliceSum(c *C) {
	for i := 0; i < c.N; i++ {
		MaxSliceSum([]int{-10})
	}
}
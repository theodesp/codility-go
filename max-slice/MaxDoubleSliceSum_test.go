package max_slice

import (
	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxDoubleSliceSum(c *C) {
	c.Assert(MaxDoubleSliceSum([]int{3, 2, 6, -1, 4, 5, -1, 2}), Equals, 17)
}

func (s *MySuite) BenchmarkMaxDoubleSliceSum(c *C) {
	for i := 0; i < c.N; i++ {
		MaxDoubleSliceSum([]int{3, 2, 6, -1, 4, 5, -1, 2})
	}
}
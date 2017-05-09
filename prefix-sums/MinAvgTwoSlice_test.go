package prefix_sums

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMinAvgTwoSlice(c *C) {
	c.Assert(MinAvgTwoSlice([]int {1, 3, 6, 4, 1, 2}), Equals, 5)
	c.Assert(MinAvgTwoSlice([]int {10, 10, -1, 2, 4, -1, 2, -1}), Equals, 5)
}

func (s *MySuite) BenchmarkMinAvgTwoSlice(c *C) {
	for i := 0; i < c.N; i++ {
		MinAvgTwoSlice([]int {1, 3, 6, 4, 1, 2})
	}
}



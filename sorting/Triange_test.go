package sorting

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMinAvgTwoSlice(c *C) {
	c.Assert(Triangle([]int{10, 2, 5, 1, 8, 20}), Equals, 1)
}

func (s *MySuite) BenchmarkMinAvgTwoSlice(c *C) {
	for i := 0; i < c.N; i++ {
		Triangle([]int{10, 2, 5, 1, 8, 20})
	}
}



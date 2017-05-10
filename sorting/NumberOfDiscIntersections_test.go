package sorting

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestNumberOfDiscIntersections(c *C) {
	c.Assert(NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}), Equals, 11)
}

func (s *MySuite) BenchmarkNumberOfDiscIntersections(c *C) {
	for i := 0; i < c.N; i++ {
		NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0})
	}
}



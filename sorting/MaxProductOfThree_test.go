package sorting

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxProductOfThree(c *C) {
	c.Assert(MaxProductOfThree([]int{1, 5, 2, 1, 4, 0}), Equals, 11)
}

func (s *MySuite) BenchmarkMaxProductOfThree(c *C) {
	for i := 0; i < c.N; i++ {
		MaxProductOfThree([]int{1, 5, 2, 1, 4, 0})
	}
}
package sorting

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxProductOfThree(c *C) {
	c.Assert(MaxProductOfThree([]int{-3, 1, 2, -2, 5, 6}), Equals, 60)
}

func (s *MySuite) BenchmarkMaxProductOfThree(c *C) {
	for i := 0; i < c.N; i++ {
		Triangle([]int{10, 2, 5, 1, 8, 20})
	}
}





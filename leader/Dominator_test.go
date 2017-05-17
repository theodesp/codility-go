package leader

import (
	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestDominator(c *C) {
	c.Assert(Dominator([]int{2, 1, 1, 1, 3}), Equals, 1)
}

func (s *MySuite) BenchmarkDominator(c *C) {
	for i := 0; i < c.N; i++ {
		Dominator([]int{2, 1, 1, 1, 3})
	}
}


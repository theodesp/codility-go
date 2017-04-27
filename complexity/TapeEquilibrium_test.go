package complexity

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestTapeEquilibrium(c *C) {
	c.Assert(TapeEquilibrium([]int{3, 1, 2, 4, 3}), Equals, 1)
}

func (s *MySuite) BenchmarkTapeEquilibrium(c *C) {
	for i := 0; i < c.N; i++ {
		TapeEquilibrium([]int{3, 1, 2, 4, 3})
	}
}

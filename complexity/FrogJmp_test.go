package complexity

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestFrogJmp(c *C) {
	c.Assert(FrogJmp(10, 85, 30), Equals, 3)
	c.Assert(FrogJmp(1, 5, 2), Equals, 2)
}

func (s *MySuite) BenchmarkFrogJmp(c *C) {
	for i := 0; i < c.N; i++ {
		FrogJmp(1, 5, 2)
	}
}

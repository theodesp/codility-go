package stacks_queues

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestNesting(c *C) {
	c.Assert(Nesting("(()(())())"), Equals, 2)
}

func (s *MySuite) BenchmarkNesting(c *C) {
	for i := 0; i < c.N; i++ {
		Nesting("(()(())())")
	}
}
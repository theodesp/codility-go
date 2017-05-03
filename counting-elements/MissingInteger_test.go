package counting_elements

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMissingInteger(c *C) {
	c.Assert(MissingInteger([]int {1, 3, 6, 4, 1, 2}), Equals, 5)
}

func (s *MySuite) BenchmarkMissingInteger(c *C) {
	for i := 0; i < c.N; i++ {
		MissingInteger([]int {1, 3, 6, 4, 1, 2})
	}
}



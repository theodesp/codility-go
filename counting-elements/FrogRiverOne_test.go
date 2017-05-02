package counting_elements

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestFrogRiverOne(c *C) {
	c.Assert(FrogRiverOne(5, []int {1, 3, 1, 4, 2, 3, 5, 4}), 6)
}

func (s *MySuite) BenchmarkFrogRiverOne(c *C) {
	for i := 0; i < c.N; i++ {
		FrogRiverOne(5, []int {1, 3, 1, 4, 2, 3, 5, 4})
	}
}


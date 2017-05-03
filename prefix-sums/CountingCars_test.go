package prefix_sums

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestFrogRiverOne(c *C) {
	c.Assert(CountingCars([]int {0, 1, 0, 1, 1}), Equals, 5)
}

func (s *MySuite) BenchmarkFrogRiverOne(c *C) {
	for i := 0; i < c.N; i++ {
		CountingCars([]int {0, 1, 0, 1, 1})
	}
}

package prefix_sums

import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestGenomicRangeQuery(c *C) {
	c.Assert(GenomicRangeQuery("AC", []int{0, 0, 1}, []int{0, 1, 1}), DeepEquals, []int{1, 1, 2})
	c.Assert(GenomicRangeQuery("A", []int{0}, []int{0}), DeepEquals, []int{1})
}

func (s *MySuite) BenchmarkGenomicRangeQuery(c *C) {
	for i := 0; i < c.N; i++ {
		c.Assert(GenomicRangeQuery("AC", []int{0, 0, 1}, []int{0, 1, 1}), DeepEquals, []int{1, 1, 2})
	}
}

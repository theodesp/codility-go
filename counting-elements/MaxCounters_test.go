package counting_elements


import (
. "gopkg.in/check.v1"
)

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxCounters(c *C) {
	c.Assert(MaxCounters(5, []int {1, 3, 6, 4, 1, 2}), DeepEquals, []int{3, 2, 2, 4, 2})
}

func (s *MySuite) BenchmarkMaxCounters(c *C) {
	for i := 0; i < c.N; i++ {
		MaxCounters(5, []int {1, 3, 6, 4, 1, 2})
	}
}



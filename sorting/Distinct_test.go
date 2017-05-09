package sorting

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestDistinct(c *C) {
	c.Assert(Distinct([]int{2, 1, 1, 2, 3, 1}), Equals, 3)
}

func (s *MySuite) BenchmarkDistinct(c *C) {
	for i := 0; i < c.N; i++ {
		Distinct([]int{2, 1, 1, 2, 3, 1})
	}
}

package leader

import (
"testing"
. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEquileader(c *C) {
	c.Assert(EquiLeader([]int{4, 3, 4, 4 ,4, 2}), Equals, 2)
}

func (s *MySuite) BenchmarkEquileader(c *C) {
	for i := 0; i < c.N; i++ {
		EquiLeader([]int{4, 3, 4, 4 ,4, 2})
	}
}


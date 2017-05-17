package max_slice

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxProfit(c *C) {
	c.Assert(MaxProfit([]int{23171, 21011, 21123, 21366, 21013, 21367}), Equals, 357)
}

func (s *MySuite) BenchmarkMaxProfit(c *C) {
	for i := 0; i < c.N; i++ {
		MaxProfit([]int{23171, 21011, 21123, 21366, 21013, 21367})
	}
}


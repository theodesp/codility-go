package prefix_sums

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCountDiv(c *C) {
	c.Assert(CountDiv(1, 1000, 11), 90)
	c.Assert(CountDiv(0, 0, 11), 1)

}

func (s *MySuite) BenchmarkCountDiv(c *C) {
	for i := 0; i < c.N; i++ {
		CountDiv(1, 1000, 11)
	}
}

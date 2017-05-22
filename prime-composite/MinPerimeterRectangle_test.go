package prime_composite

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMinPerimeterRectangle(c *C) {
	c.Assert(MinPerimeterRectangle(30), Equals, 22)
}

func (s *MySuite) BenchmarkMinPerimeterRectangle(c *C) {
	for i := 0; i < c.N; i++ {
		MinPerimeterRectangle(30)
	}
}

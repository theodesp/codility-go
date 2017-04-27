package complexity

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPermMissingElem(c *C) {
	var solutions = [...]int{4, 6, 3}

	c.Assert(PermMissingElem([]int{2, 3, 1, 5}), Equals, solutions[0])
	c.Assert(PermMissingElem([]int{1, 2, 3, 4, 5, 7}), Equals, solutions[1])
}

func (s *MySuite) BenchmarkPermMissingElem(c *C) {
	for i := 0; i < c.N; i++ {
		PermMissingElem([]int{1, 2, 3, 4, 5, 7})
	}
}


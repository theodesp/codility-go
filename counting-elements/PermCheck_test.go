package counting_elements

import (
"testing"
. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPermCheck(c *C) {
	var solutions = [...]int{0, 1}
	c.Assert(PermCheck([]int{4, 1, 3, 2}), Equals, 1)
	c.Assert(PermCheck([]int{1, 4, 1}), Equals, 0)
}



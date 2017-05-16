package stacks_queues

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestBrackets(c *C) {
	c.Assert(Brackets("{[()()]}") ,Equals, 1)
	c.Assert(Brackets("([)()]") ,Equals, 0)
	c.Assert(Brackets(")(") ,Equals, 0)

}

func (s *MySuite) BenchmarkBrackets(c *C) {
	for i := 0; i < c.N; i++ {
	}
}

package iterations

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSolution(c *C) {
	var solutions = [...]int {0, 0, 0, 0, 1, 0, 2, 5, 0}

	c.Assert(Solution(1), Equals, solutions[0])
	c.Assert(Solution(2), Equals, solutions[1])
	c.Assert(Solution(3), Equals, solutions[2])
	c.Assert(Solution(4), Equals, solutions[3])
	c.Assert(Solution(5), Equals, solutions[4])
	c.Assert(Solution(6), Equals, solutions[5])
	c.Assert(Solution(200), Equals, solutions[6])
	c.Assert(Solution(1041), Equals, solutions[7])
	c.Assert(Solution(1<<31 - 2), Equals, solutions[8])
}

func (s *MySuite) BenchmarkSolution(c *C) {
	for i := 0; i < c.N; i++ {
		Solution(1<<31)
	}
}

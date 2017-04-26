package arrays

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCyclicRotation(c *C) {
	c.Assert(CyclicRotate([]int{0 -9}, 2), DeepEquals, []int{-9})
	c.Assert(CyclicRotate([]int{0 -9}, 1), DeepEquals, []int{-9})
	c.Assert(CyclicRotate([]int{1, 1, 2, 3, 5}, 42), DeepEquals, []int{3, 5, 1, 1, 2})
	c.Assert(CyclicRotate([]int{3, 8, 9, 7, 6}, 3), DeepEquals, []int{9, 7, 6, 3, 8})
}

func (s *MySuite) BenchmarkCyclicRotation(c *C) {
	for i := 0; i < c.N; i++ {
		CyclicRotate([]int{3, 8, 9, 7, 6}, 3)
	}
}

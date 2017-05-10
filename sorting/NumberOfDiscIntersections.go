package sorting

import (
	"sort"
)

/*
We draw N discs on a plane. The discs are numbered from 0 to N − 1.
A zero-indexed array A of N non-negative integers, specifying the radiuses of the discs, is given.
The J-th disc is drawn with its center at (J, 0) and radius A[J].

We say that the J-th disc and K-th disc intersect if J ≠ K and the J-th and K-th discs
have at least one common point (assuming that the discs contain their borders).

The figure below shows discs drawn for N = 6 and A as follows:

  A[0] = 1
  A[1] = 5
  A[2] = 2
  A[3] = 1
  A[4] = 4
  A[5] = 0


There are eleven (unordered) pairs of discs that intersect, namely:

discs 1 and 4 intersect, and both intersect with all the other discs;
disc 2 also intersects with discs 0 and 3.
Write a function:

func Solution(A []int) int
that, given an array A describing N discs as explained above, returns the number of
(unordered) pairs of intersecting discs. The function should return −1 if the number of intersecting pairs exceeds 10,000,000.

Given array A shown above, the function should return 11, as explained above.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [0..2,147,483,647].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
 */

// Map circles to points on the left and on the right of the line segment -x, x and sort them
// Then for every left point do a binary search for the index of the other point. Count the number of the
// elements before that index and subtract i+1 to remove duplicate counts
//


func NumberOfDiscIntersections(A []int) int {
	lefts := make([]int, len(A))
	rights := make([]int, len(A))
	result := 0

	for i := 0; i < len(A); i += 1 {
		lefts[i] = i - A[i]
		rights[i] = i + A[i]
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	for i := 0; i < len(A); i += 1 {
		end := rights[i]

		// Binary search for index of element of the rightmost value less than to the interval-end
		count := sort.Search(len(lefts), func(i int) bool {
			return lefts[i] > end
		})

		count -= (i + 1)
		result += count

		if result > 10000000 {
			return -1
		}

	}

	return result
}


//package solution

//// you can also use imports, for example:
//// import "fmt"
//// import "os"
//
//import "sort"
//
//// you can write to stdout for debugging purposes, e.g.
//// fmt.Println("this is a debug message")
//
//func Solution(A []int) int {
//	pairs := [][]int{}
//	starts := make([]int, len(A))
//	result := 0
//
//	for i := 0; i < len(A); i += 1 {
//		pairs = append(pairs, []int{i - A[i], i + A[i]})
//	}
//
//	sort.Slice(pairs[:], func(i, j int) bool {
//		for x := range pairs[i] {
//			if pairs[i][x] == pairs[j][x] {
//				continue
//			}
//			return pairs[i][x] < pairs[j][x]
//		}
//		return false
//	})
//
//	for i := 0; i < len(A); i += 1 {
//		starts[i] = pairs[i][0]
//	}
//
//	for i := 0; i < len(A); i += 1 {
//		end := pairs[i][1]
//
//		// Binary search for index of element of the rightmost value less than to the interval-end
//		count := sort.Search(len(starts), func(i int) bool {
//			return starts[i] > end
//		})
//
//		count -= (i + 1)
//		result += count
//		if result > 10000000 {
//			return -1
//		}
//
//	}
//
//	return result
//}
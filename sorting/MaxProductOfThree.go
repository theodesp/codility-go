package sorting

import "math"

/*
A non-empty zero-indexed array A consisting of N integers is given.
The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
contains the following example triplets:

(0, 1, 2), product is −3 * 1 * 2 = −6
(1, 2, 4), product is 1 * 2 * 5 = 10
(2, 4, 5), product is 2 * 5 * 6 = 60
Your goal is to find the maximal product of any triplet.

Write a function:

func Solution(A []int) int
that, given a non-empty zero-indexed array A, returns the value of the maximal product of any triplet.

For example, given array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Assume that:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−1,000..1,000].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(1), beyond input storage
(not counting the storage required for input arguments).
 */

// Find the largest element a, the second largest element b and the 3rd largest element c. Find also the smallest element d
// and the second smallest e element. Return max(a*b*c, a*d*e)
// For example if we had the 2 smallest elements to be negative it would have a chance to be the answer as they can produce a
// positive nuber
func MaxProductOfThree(A []int) int {
	arrayLen := len(A)

	if arrayLen < 3 {
		// Early exit
		return 0
	}

	a, b, c := -1 << 31, -1 << 31, -1 << 31
	d, e := 1 << 31 - 1, 1 << 31 - 1

	for i := 0; i < arrayLen; i += 1 {

		// Update maximums
		if A[i] > a {
			c = b
			b = a
			a = A[i]
		} else if A[i] > b {
			c = b
			b = A[i]
		} else if A[i] > c {
			c = A[i]
		}

		// Update minimums
		if A[i] < d {
			e = d
			d = A[i]
		} else if A[i] < e {
			e = A[i]
		}
	}

	return int(math.Max(float64(a) * float64(b) * float64(c), float64(a) * float64(d) * float64(e)));
}

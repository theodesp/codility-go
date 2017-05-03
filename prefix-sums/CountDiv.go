package prefix_sums

import "math"

/*
Write a function:

func Solution(A int, B int, K int) int

that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

{ i : A ≤ i ≤ B, i mod K = 0 }
For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by
2 within the range [6..11], namely 6, 8 and 10.

Assume that:

A and B are integers within the range [0..2,000,000,000];
K is an integer within the range [1..2,000,000,000];
A ≤ B.

Complexity:

expected worst-case time complexity is O(1);
expected worst-case space complexity is O(1).
 */

// B/K is the largest multiple of K. Subtract that from largest A-1/K if A > 0 to find the required numbers
func CountDiv(A int, B int, K int) int {
	addOne := 0

	if A % K == 0 {
		addOne = 1
	}

	return int(math.Floor(float64(B / K))) - int(math.Floor(float64((A) / K))) + addOne
}
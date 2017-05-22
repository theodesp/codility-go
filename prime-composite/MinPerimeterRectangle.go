package prime_composite

import (
	"math"
)

/*
An integer N is given, representing the area of some rectangle.

The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

The goal is to find the minimal perimeter of any rectangle whose area equals N.
The sides of this rectangle should be only integers.

For example, given integer N = 30, rectangles of area 30 are:

(1, 30), with a perimeter of 62,
(2, 15), with a perimeter of 34,
(3, 10), with a perimeter of 26,
(5, 6), with a perimeter of 22.
Write a function:

func Solution(N int) int
that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

For example, given an integer N = 30, the function should return 22, as explained above.

Assume that:

N is an integer within the range [1..1,000,000,000].
Complexity:

expected worst-case time complexity is O(sqrt(N));
expected worst-case space complexity is O(1).
 */

// FInd the devisors of N lets say (A1,A2...AN) and go for each on of them and calc min 2 * (AN * BN)
func MinPerimeterRectangle(N int) int {
	result := 1 << 32 - 1
	i := 1
	pairs := make(map[int]int)

	for i * i < N {
		if N % i == 0 {
			pairs[i] = N / i
		}

		i += 1
	}

	if i * i == N {
		pairs[i] = i
	}

	for j, value := range (pairs) {
		perimeter := 2 * (j + value)
		result = int(math.Min(float64(result), float64(perimeter)))
	}

	return result
}


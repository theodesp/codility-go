package iterations

/*
 A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones
 at both ends in the binary representation of N.

 For example, number 9 has binary representation 1001 and contains a binary gap of length 2.
 The number 529 has binary representation 1000010001 and contains two binary gaps: one of
 length 4 and one of length 3. The number 20 has binary representation 10100 and contains one
 binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps.

 Write a function:

 func Solution(N int) int

 that, given a positive integer N, returns the length of its longest binary gap.
 The function should return 0 if N doesn't contain a binary gap.

 For example, given N = 1041 the function should return 5, because N has binary
 representation 10000010001 and so its longest binary gap is of length 5.
 */


// See https://graphics.stanford.edu/~seander/bithacks.html#ZerosOnRightLinear
func countTrailingBits(N uint) uint {
	var _lookup = [...]uint{
		32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4,
		7, 17, 0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5,
		20, 8, 19, 18 }

	return _lookup[(N & -N) % 37]
}

/*
 * Iterate through the bits and find the longest
 * substring of zeros in O(L) where L is the length of the binary string => floor(log(n+1))
 */
func Solution(N int) (maxLength int) {
	maxLength = 0
	currentLength := 0
	trailing := countTrailingBits(uint(N))

	if trailing > 0 {
		N = int(N >> trailing)
	}

	for N > 0 {
		// N % 2 gets the bit at the rightmost part and if its 1 then reset the counter
		if N % 2 == 1 {
			currentLength = 0
		} else {
			currentLength += 1
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}

		N >>= 1
	}

	return;
}

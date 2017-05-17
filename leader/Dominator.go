package leader

import (
	"math"
)

/*
A zero-indexed array A consisting of N integers is given. The dominator of array A is
the value that occurs in more than half of the elements of A.

For example, consider array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those
with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

Write a function

func Solution(A []int) int
that, given a zero-indexed array A consisting of N integers, returns index of any element of
array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
For example, given array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
the function may return 0, 2, 4, 6 or 7, as explained above.

Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).
 */

func FindIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func Dominator(A []int) int {
	// O(N)
	arrayLen := len(A)

	if arrayLen == 1 {
		return 0
	}

	l := NewIntStack(arrayLen)
	candidate := -1
	leader := -1
	count := 0

	for _, value := range (A) {
		//O(n)
		if l.size == 0 {
			l.Push(value)
		} else {
			if value != l.Front() {
				l.Pop()
			} else {
				l.Push(value)
			}
		}
	}

	if l.size > 0 {
		candidate = l.Front()
	} else {
		return -1
	}

	for _, value := range (A) {
		if value == candidate {
			count += 1
		}
	}

	if count > int(math.Floor(float64(arrayLen) / 2.0)) {
		leader = candidate
	} else {
		return -1
	}

	return FindIndex(arrayLen, func(i int) bool { return A[i] == leader })
}

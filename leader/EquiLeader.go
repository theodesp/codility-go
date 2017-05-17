package leader

import (
	"math"
)

/*
A non-empty zero-indexed array A consisting of N integers is given.

The leader of this array is the value that occurs in more than half of the elements of A.

An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1]
have leaders of the same value.

For example, given array A such that:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
we can find two equi leaders:

0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
The goal is to count the number of equi leaders.

Write a function:

func Solution(A []int) int
that, given a non-empty zero-indexed array A consisting of N integers, returns the number of equi leaders.

For example, given:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
the function should return 2, as explained above.

Assume that:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
 */
/*
	8x
	6  6x 8x
	6x     6  6  6  6x
	4 4x  6  6  6  6  6
	----------------------
	|4 |6 |6 |6 |6 |8 |8 |
	----------------------
*/

func Leader(A []int) int {
	// O(N)
	arrayLen := len(A)

	if arrayLen == 1 {
		return A[0]
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
	}

	for _, value := range (A) {
		if value == candidate {
			count += 1
		}
	}

	if count > int(math.Floor(float64(arrayLen) / 2.0)) {
		leader = candidate
	}

	return leader

}

// We utilize the Leader implementation internals to maintain a list of subleaders in a map
func EquiLeader(A []int) int {
	leadersCount := 0
	arrayLen := len(A)
	l := NewIntStack(arrayLen)
	candidate := -1
	leader := -1
	count := 0
	leftLeadersCount := 0
	leftSequenceLength := 0
	rightSequenceLength := 0

	if arrayLen == 1 {
		return 0
	}

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
		return 0
	}

	for _, value := range (A) {
		if value == candidate {
			count += 1
		}
	}

	if count > int(math.Floor(float64(arrayLen) / 2.0)) {
		leader = candidate
	} else {
		return 0
	}

	for i, value := range (A) {
		if value == leader {
			// The key here is to compare the current leader with the current value
			leftLeadersCount += 1
		}

		/* validate left sequence */
		leftSequenceLength = i + 1;

		if leftLeadersCount > int(math.Floor(float64(leftSequenceLength) / 2.0)) {
			/* validate right sequence */
			rightSequenceLength = arrayLen - leftSequenceLength;

			// Here we check if the remaining leaders count on the left side is more than have of the remaining on the right side
			if count - leftLeadersCount > int(math.Floor(float64(rightSequenceLength) / 2.0)) {
				/* both sequences have valid leaders of the same value */
				leadersCount += 1
			}
		}
	}

	return leadersCount
}

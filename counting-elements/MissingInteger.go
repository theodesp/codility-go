package counting_elements

/*
Write a function:

func Solution(A []int) int
that, given a non-empty zero-indexed array A of N integers, returns the minimal positive
integer (greater than 0) that does not occur in A.

For example, given:

  A[0] = 1
  A[1] = 3
  A[2] = 6
  A[3] = 4
  A[4] = 1
  A[5] = 2
the function should return 5.

Assume that:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the
storage required for input arguments).
 */

// Iterate over the elements and find the first one 1 and then collect all positive elements in
// map. Then iterate over the elements and see which one is missing
func MissingInteger(A []int) int {
	init := 1
	seen := make(map[int]bool)

	for _, value := range(A) {
		if value > 0 && !seen[value] {
			seen[value] = true
		}
	}

	for i := init; i <= len(seen); i+=1 {
		if !seen[init] {
			return init
		}
		init += 1
	}

	return init
}

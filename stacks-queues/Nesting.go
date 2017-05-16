package stacks_queues


/*
A string S consisting of N characters is called properly nested if:

S is empty;
S has the form "(U)" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

func Solution(S string) int
that, given a string S consisting of N characters, returns 1 if string S is
properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())",
the function should return 0, as explained above.

Assume that:

N is an integer within the range [0..1,000,000];
string S consists only of the characters "(" and/or ")".
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1) (not counting the storage required for input arguments).

 */

func Nesting(S string) int {
	l := NewStringStack(len(S))

	if len(S) == 0 {
		return 1
	}

	if len(S) % 2 != 0 {
		return 0
	}

	for _, value := range (S) {
		val := string(value)
		if val == "(" {
			l.Push(val)
		} else if val == ")" {
			if l.size == 0 {
				return 0
			}

			head := l.Front()
			if head == "(" {
				l.Pop()
			} else {
				return 0
			}
		}
	}

	if l.size == 0 {
		return 1
	} else {
		return 0
	}
}

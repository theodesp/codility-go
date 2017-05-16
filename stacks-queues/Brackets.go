package stacks_queues

/*
A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

func Solution(S string) int
that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Assume that:

N is an integer within the range [0..200,000];
string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).

 */

func Brackets(S string) int {
	l := NewStringStack(len(S))
	pairs := make(map[string]string)
	pairs["{"] = "}"
	pairs["["] = "]"
	pairs["("] = ")"

	if len(S) == 0 {
		return 1
	}

	if len(S) % 2 != 0 {
		return 0
	}

	for _, value := range (S) {
		val := string(value)
		if val == "(" || val == "{" || val == "[" {
			l.Push(val)
		} else if val == ")" || val == "}" || val == "]" {
			if l.size == 0 {
				return 0
			}

			head := l.Front()
			if pairs[head] == val {
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
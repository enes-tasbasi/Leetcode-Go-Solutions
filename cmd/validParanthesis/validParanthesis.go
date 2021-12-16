package validparanthesis

import (
	"container/list"
	"fmt"
)

func main() {
	str := "[({})[]]"

	fmt.Printf("%v\n", isValid(str))
}

func isValid(s string) bool {
	inStr := []string{"(", "{", "["}
	outStr := []string{")", "}", "]"}
	stack := list.New()

	for i := 0; i < len(s); i++ {
		currChar := string(s[i])

		inIndex := -1
		for inI := 0; inI < len(inStr); inI++ {
			if currChar == inStr[inI] {
				inIndex = inI
				break
			}
		}

		if inIndex != -1 {
			stack.PushFront(inIndex)
			continue
		}

		outIndex := -1
		for outI := 0; outI < len(outStr); outI++ {
			if currChar == outStr[outI] {
				outIndex = outI
				break
			}
		}

		if outIndex != -1 {
			if stack.Len() == 0 {
				return false
			}
			if stack.Front().Value != outIndex {
				return false
			}

			stack.Remove(stack.Front())
		}
	}

	return stack.Len() == 0
}

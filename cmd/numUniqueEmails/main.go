package main

import (
	"fmt"
	"strings"
)

func main() {
	t := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}

	fmt.Println(numUniqueEmails(t))
}

// Every valid email consists of a local name and a domain name, separated by the '@' sign.
// Besides lowercase letters, the email may contain one or more '.' or '+'.
func numUniqueEmails(emails []string) int {
	res := make(map[string]bool)

	for _, email := range emails {
		s := strings.Split(email, "@")
		local := s[0]

		local = strings.ReplaceAll(local, ".", "")

		if strings.Contains(local, "+") {
			s2 := strings.Split(local, "+")

			res[fmt.Sprintf("%s@%s", s2[0], s[1])] = true
		} else {
			res[fmt.Sprintf("%s@%s", local, s[1])] = true
		}
	}

	return len(res)
}

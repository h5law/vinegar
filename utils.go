package vinegar

import "strings"

// formatKeyWord removes any non-latin characters, removes all spaces and
// converts the given keyword to lowercase.
func formatKeyword(keyword string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return -1
		}
		if r > 'z' || r < 'a' {
			if r >= 'A' && r <= 'Z' {
				r += 32
			} else {
				return -1
			}
		}
		return r
	}, keyword)
}

// formatTableKeyword calls formatKeyword then removes any duplicated and
// ensures the keyword is at most 26 characters long.
func formatTableKeyword(keyword string) string {
	keyword = formatKeyword(keyword)
	keyword = removeDuplicates(keyword)
	if len(keyword) > 26 {
		return keyword[:26]
	}
	return keyword
}

// removeDuplicates removes any duplicates from the provided keyword.
func removeDuplicates(s string) string {
	str := strings.Builder{}
	seen := make([]int, 26)
	for _, c := range s {
		if seen[c%26] == 0 {
			str.WriteRune(c)
		} else {
			seen[c%26] = 1
		}
	}
	return str.String()
}

// formatEncryptionKeyword ensures the keyword is formatted with formatKeyword
// then is repeated (and spliced if necessary) to be the same length as the message.
func formatEncryptionKeyword(keyword, message string) string {
	keyword = formatKeyword(keyword)
	k, m := len(keyword), len(message)
	if k == m {
		return keyword
	}
	if k > m {
		return keyword[:m]
	}
	str := strings.Builder{}
	str.WriteString(keyword)
	for str.Len() != m {
		if m-str.Len() >= k {
			str.WriteString(keyword)
		} else {
			str.WriteString(keyword[:m-str.Len()])
		}
	}
	return str.String()
}

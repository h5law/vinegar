package vinegar

import "strings"

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

func formatTableKeyword(keyword string) string {
	keyword = formatKeyword(keyword)
	keyword = removeDuplicates(keyword)
	if len(keyword) > 26 {
		return keyword[:26]
	}
	return keyword
}

func removeDuplicates(s string) string {
	str := strings.Builder{}
	seen := make(map[rune]int)
	for _, c := range s {
		if _, ok := seen[c]; ok {
			continue
		}
		str.WriteRune(c)
		seen[c]++
	}
	return str.String()
}

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

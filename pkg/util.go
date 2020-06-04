package pkg

import "strings"

func LeftAdjust(s string, a string) string {
	ps := strings.Split(s, "\n")
	res := ""
	for i, p := range ps {
		res += a + p
		if i != len(ps)-1 {
			res += "\n"
		}
	}
	return res
}

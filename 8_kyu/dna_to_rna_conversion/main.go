package kata

import "strings"

func DNAtoRNA(dna string) string {
	result := strings.Builder{}
	for _, c := range dna {
		if c == 'T' {
			result.WriteRune('U')
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}

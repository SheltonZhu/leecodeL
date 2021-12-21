package _2_generate_parentheses

import "bytes"

func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, &bytes.Buffer{}, n, 0, 0)
	return result
}

func backtrack(result *[]string, b *bytes.Buffer, n, left, right int) {
	if b.Len() == 2*n {
		*result = append(*result, b.String())
	} else {
		tmp := b.String()
		if right < n && left > right {
			b.WriteString(")")
			backtrack(result, b, n, left, right+1)
		}
		if left < n {
			b.Reset()
			b.WriteString(tmp)
			b.WriteString("(")
			backtrack(result, b, n, left+1, right)
		}
	}
}

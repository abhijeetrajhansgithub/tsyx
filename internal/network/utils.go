package net

import (
	"strings"
)


func splitFieldName(name string) string {
	var b strings.Builder

	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			b.WriteByte(' ')
		}
		b.WriteRune(r)
	}

	return b.String()
}
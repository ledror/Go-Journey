package comma

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	if len(s) == 0 {
		return ""
	}
	b := new(bytes.Buffer)
	integerStart := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		integerStart++
	}

	var fraction string
	integerEnd := strings.Index(s, ".")
	if integerEnd == -1 {
		integerEnd = len(s)
	} else {
		fraction = s[integerEnd:]
	}
	integer := s[integerStart:integerEnd]

	skip := len(integer) % 3
	if skip == 0 {
		skip = 3
	}
	b.WriteString(integer[:skip])
	for i := skip; i < len(integer); i += 3 {
		b.WriteByte(',')
		b.WriteString(integer[i : i+3])
	}
	if fraction == "" {
		return b.String()
	}

	b.WriteByte('.')
	fraction = fraction[1:]
	skip = len(fraction) % 3
	if skip == 0 {
		skip = 3
	}
	// decimal commas are from left to right
	for i := 0; i < len(fraction)-skip; i += 3 {
		b.WriteString(fraction[i : i+3])
		b.WriteByte(',')
	}
	b.WriteString(fraction[len(fraction)-skip:])

	return b.String()
}

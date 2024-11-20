package query

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var safeSet = [utf8.RuneSelf]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
	'G': true,
	'H': true,
	'I': true,
	'J': true,
	'K': true,
	'L': true,
	'M': true,
	'N': true,
	'O': true,
	'P': true,
	'Q': true,
	'R': true,
	'S': true,
	'T': true,
	'U': true,
	'V': true,
	'W': true,
	'X': true,
	'Y': true,
	'Z': true,
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
	'@': true,
	'-': true,
	'.': true,
	'_': true,
}

func needsQuoting(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); {
		b := s[i]
		if b < utf8.RuneSelf {
			// Quote anything except a backslash that would need quoting in a
			// JSON string, as well as space and '='
			if b != '\\' && !safeSet[b] {
				return true
			}
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError || unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}
		i += size
	}
	return false
}

type raw struct {
	str string
}

func (r raw) String() string {
	return r.str
}

type not struct {
	query Query
}

func (n not) String() string {
	if n.query == nil {
		return ""
	}
	str := n.query.String()
	if str == "" {
		return ""
	}
	return "-" + str
}

type Query interface {
	String() string
}

type field struct {
	key   string
	value string
}

func (f field) String() string {
	if f.key == "" {
		return f.value
	}
	var b strings.Builder
	b.WriteString(f.key)
	b.WriteByte(':')
	if needsQuoting(f.value) {
		b.WriteString(strconv.Quote(f.value))
	} else {
		b.WriteString(f.value)
	}
	return b.String()
}

func queriesToString(qs []Query, sep string) string {
	if len(qs) == 0 {
		return ""
	}
	if len(qs) == 1 {
		return qs[0].String()
	}
	var b strings.Builder
	b.WriteByte('(')
	var firstIndex int
	for i, q := range qs {
		str := q.String()
		if str != "" {
			b.WriteString(str)
			firstIndex = i + 1
			break
		}
	}
	if firstIndex == 0 {
		return ""
	}

	for _, q := range qs[firstIndex:] {
		str := q.String()
		if str == "" {
			continue
		}
		b.WriteString(sep)
		b.WriteString(str)
	}
	b.WriteByte(')')
	return b.String()
}

type and []Query

func (a and) String() string {
	return queriesToString(a, " AND ")
}

type or []Query

func (a or) String() string {
	return queriesToString(a, " OR ")
}

func F(key, value string) Query {
	return field{key, value}
}

func Or(queries ...Query) Query {
	return or(queries)
}

func And(queries ...Query) Query {
	return and(queries)
}

func Not(query Query) Query {
	return not{query}
}

func Raw(str string) Query {
	return raw{str}
}

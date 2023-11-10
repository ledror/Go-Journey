package main

import (
	"strings"
)

func variadicJoin(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

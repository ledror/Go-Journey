package comma

func CommaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return CommaRec(s[:n-3]) + "," + s[n-3:]
}

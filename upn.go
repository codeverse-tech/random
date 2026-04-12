func upn(s string) []string {
	f := strings.Fields(s)

	for i := 0; i < len(f); i++ {
		if f[i] == "(up," {
			next := f[i+1]
			next = strings.TrimSuffix(next, ")")
			conv, _ := strconv.Atoi(next)

			for j := i - conv; j < i; j++ {
				f[j] = strings.ToUpper(f[j])
			}
			f = append(f[:i], f[i+2:]...)
			i--
		}
	}
	return f
}
package unique

func unique(strs []string) []string {
	i := 0 // index of last written string
	for _, s := range strs {
		if strs[i] == s {
			continue
		}
		i++
		strs[i] = s
	}
	return strs[:i+1]
}

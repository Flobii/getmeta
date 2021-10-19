package stringsplus

func GetMaxLength(strings []string) int {
	max := -1
	for _, s := range strings {
		if len(s) < max {
			continue
		}
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func GetLongest(strings []string) string {
	max := -1
	idx := -1
	for i, s := range strings {
		if len(s) < max {
			continue
		}
		if len(s) > max {
			max = len(s)
			idx = i
		}
	}
	return strings[idx]
}


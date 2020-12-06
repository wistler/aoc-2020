package internal

// Contains searches for needle in haystack. If found, returns index.
func Contains(haystack []string, needle string) (bool, int) {
	for i, r := range haystack {
		if r == needle {
			return true, i
		}
	}
	return false, -1
}

// RemoveIndex pops element at given index of slice
func RemoveIndex(from *[]string, index int) {
	f := *from
	f[index] = f[len(f)-1]
	f[len(f)-1] = ""
	f = f[:len(f)-1]
	*from = f
}

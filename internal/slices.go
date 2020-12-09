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

// ContainsNumber for int[]
func ContainsNumber(haystack []int, needle int) (bool, int) {
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

// EqualI checks if two int slices are equal
func EqualI(p []int, b []int) bool {
	if len(p) != len(b) {
		return false
	}
	for i, v := range p {
		if v != b[i] {
			return false
		}
	}
	return true
}

// EqualS checks if two int slices are equal
func EqualS(p []string, b []string) bool {
	if len(p) != len(b) {
		return false
	}
	for i, v := range p {
		if v != b[i] {
			return false
		}
	}
	return true
}

// SumOf sum of
func SumOf(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Min min
func Min(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if min > n {
			min = n
		}
	}
	return min
}

// Max max
func Max(numbers []int) int {
	max := numbers[0]
	for _, n := range numbers {
		if max < n {
			max = n
		}
	}
	return max
}

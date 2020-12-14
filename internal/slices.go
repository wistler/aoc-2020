package internal

import "hash/fnv"

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

// ContainsRune for rune[]
func ContainsRune(haystack []rune, needle rune) (bool, int) {
	for i, r := range haystack {
		if r == needle {
			return true, i
		}
	}
	return false, -1
}

// ContainsByte for byte[]
func ContainsByte(haystack []byte, needle byte) (bool, int) {
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

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Interface is a type for Equal check; inspired from go sort package
type Interface interface {
	len() int
	getHash(i int) uint32
}

type stringSlice []string
type intSlice []int

func (p stringSlice) len() int             { return len(p) }
func (p stringSlice) getHash(i int) uint32 { return hash(p[i]) }
func (p intSlice) len() int                { return len(p) }
func (p intSlice) getHash(i int) uint32    { return uint32(p[i]) }

// Equal checks if two data slices are equal
func Equal(a Interface, b Interface) bool {
	if a.len() != b.len() {
		return false
	}
	for i := 0; i < a.len(); i++ {
		if a.getHash(i) != b.getHash(i) {
			return false
		}
	}
	return true
}

// IntsAreEqual checks if two int slices are equal
func IntsAreEqual(a []int, b []int) bool {
	return Equal(intSlice(a), intSlice(b))
}

// StringsAreEqual checks if two int slices are equal
func StringsAreEqual(a []string, b []string) bool {
	return Equal(stringSlice(a), stringSlice(b))
}

// ints is a slice of ints
type ints []int

// Sum returns sum of all elements in slice
func (p ints) Sum() int {
	return SumOf(p)
}

// Min returns lowest number in slice
func (p ints) Min() int {
	return Min(p)
}

// Max returns highest number in slice
func (p ints) Max() int {
	return Max(p)
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

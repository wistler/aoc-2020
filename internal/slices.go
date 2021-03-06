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
	Len() int
	GetHash(i int) uint32
}

type stringSlice []string
type intSlice []int
type boolSlice []bool

func (p stringSlice) Len() int             { return len(p) }
func (p stringSlice) GetHash(i int) uint32 { return hash(p[i]) }
func (p intSlice) Len() int                { return len(p) }
func (p intSlice) GetHash(i int) uint32    { return uint32(p[i]) }
func (p boolSlice) Len() int               { return len(p) }
func (p boolSlice) GetHash(i int) uint32 {
	if p[i] {
		return 1
	}
	return 0
}

// Equal checks if two data slices are equal
func Equal(a Interface, b Interface) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.GetHash(i) != b.GetHash(i) {
			return false
		}
	}
	return true
}

// IntsAreEqual checks if two int slices are equal
func IntsAreEqual(a []int, b []int) bool {
	return Equal(intSlice(a), intSlice(b))
}

// BoolsAreEqual checks if two bool slices are equal
func BoolsAreEqual(a []bool, b []bool) bool {
	return Equal(boolSlice(a), boolSlice(b))
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

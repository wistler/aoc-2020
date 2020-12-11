package internal

import "testing"

func TestIntsAreEqual(t *testing.T) {
	testCases := []struct {
		desc   string
		input  [][]int
		output bool
	}{
		{
			desc: "Ints of same content",
			input: [][]int{
				{1, 2, 3},
				{1, 2, 3},
			},
			output: true,
		},
		{
			desc: "Ints of same content (longer)",
			input: [][]int{
				{-1, -2, 3, 0, 0, 0, 0, 0, 4},
				{-1, -2, 3, 0, 0, 0, 0, 0, 4},
			},
			output: true,
		},
		{
			desc: "Empty slices are equal",
			input: [][]int{
				{},
				{},
			},
			output: true,
		},
		{
			desc: "1-elemnt slices are equal",
			input: [][]int{
				{2048},
				{2048},
			},
			output: true,
		},
		{
			desc: "1-elemnt slices with different content",
			input: [][]int{
				{1},
				{2},
			},
			output: false,
		},
		{
			desc: "Slices of different lengths",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4},
			},
			output: false,
		},
		{
			desc: "One of the slices is empty",
			input: [][]int{
				{},
				{1, 2, 3, 4},
			},
			output: false,
		},
	}
	failed := false
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := IntsAreEqual(tC.input[0], tC.input[1])
			if got != tC.output {
				failed = true
				t.Logf("For inputs: %v, %v\nWanted: %v, but got: %v",
					tC.input[0], tC.input[1], tC.output, got)
			}
		})
	}
	if failed {
		t.FailNow()
	}
}

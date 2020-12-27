package day20

import (
	"fmt"

	"github.com/wistler/aoc-2020/internal"
)

type pixel struct {
	data bool
}

type pixels []pixel
type image []pixels

func (p pixels) Len() int {
	return len(p)
}

func (p pixels) GetHash(i int) uint32 {
	if p[i].data {
		return 1
	}
	return 0
}

func pixelsMatch(a pixels, b pixels) bool {
	return internal.Equal(pixels(a), pixels(b))
}

type tileState struct {
	rotated int
	flipped bool
}

type tile struct {
	id       int
	rawData  image
	faceDown bool
	rotation int
	links    [4]int
}

func toStr(bit pixel) string {
	if bit.data {
		return "#"
	}
	return "."
}

func toString(row pixels) string {
	str := ""
	for _, b := range row {
		str += toStr(b)
	}
	return str
}

func (t tile) String() string {
	str := ""
	for _, bor := range t.rawData {
		str += fmt.Sprintf("%s\n", toString(bor[:]))
	}
	f := "🔺"
	if t.faceDown {
		f = "🔻"
	}
	LINK := []string{"⬆", "➡", "⬇", "⬅"}
	links := ""
	for l := 0; l < 4; l++ {
		links += fmt.Sprintf("%3s %4d\n", LINK[l], t.links[l])
	}
	return fmt.Sprintf("\n%d(%v%v):\n%s%s", t.id, t.rotation, f, str, links)
}

func (t *tile) getBorders() [4]pixels {
	return getBorders(t.rawData)
}

func (t *tile) getData() image {
	return removeBorder(t.rawData)
}

func (t *tile) getLinkCount() int {
	count := 0
	for _, l := range t.links {
		if l != 0 {
			count++
		}
	}
	return count
}

func (t *tile) isCorner() bool {
	return t.getLinkCount() == 2
}

func getBorders(d image) [4]pixels {
	if len(d) != len(d[0]) {
		panic("Expected square matrix")
	}

	S := len(d) - 1
	var borders [4]pixels
	borders[0] = d[0]
	borders[2] = d[S]
	borders[1] = make(pixels, 10)
	borders[3] = make(pixels, 10)
	for i := 0; i <= S; i++ {
		borders[1][i] = d[i][S]
		borders[3][S-i] = d[i][0]
	}
	return borders
}

func removeBorder(d image) image {
	if len(d) != len(d[0]) {
		panic("Expected square matrix")
	}

	S := len(d) - 2
	data := make(image, S)
	for i := 0; i < S; i++ {
		data[i] = make(pixels, S)
		for j := 0; j < S; j++ {
			data[i][j] = d[i+1][j+1]
		}
	}
	return data
}

// returns a new tile instance with applied transform
func (t tile) transform(state tileState) tile {
	f := false
	if state.flipped {
		t.faceDown = !t.faceDown
		f = true
	}

	r := (4 + state.rotated) % 4
	if t.faceDown {
		r = (4 - state.rotated) % 4
	}
	t.rotation = (t.rotation + r) % 4
	if f {
		t.rawData = flip(t.rawData)
		t.links = applyState(t.links, tileState{0, true})
	}
	if r != 0 {
		t.rawData = rotate(t.rawData, r)
		t.links = applyState(t.links, tileState{r, false}) // FIXME: Bug: Needs to account for 🔺🔻 ?
	}
	return t
}

func (t *tile) getState() tileState {
	return tileState{t.rotation, t.faceDown}
}

func rotate(data image, rotate int) image {
	if len(data) != len(data[0]) {
		panic("data must be a square matrix")
	}

	S := len(data)
	r := (4 + rotate) % 4

	read := make(image, S)
	for i := 0; i < S; i++ {
		read[i] = make(pixels, S)
		for j := 0; j < S; j++ {
			switch r {
			case 0:
				read[i][j] = data[i][j]
			case 1:
				read[i][j] = data[j][S-1-i]
			case 2:
				read[i][j] = data[S-1-i][S-1-j]
			case 3:
				read[i][j] = data[S-1-j][i]
			}
		}
	}
	return read
}

func flip(data image) image {
	if len(data) != len(data[0]) {
		panic("data must be a square matrix")
	}

	S := len(data)

	fli := make(image, S)
	for i := 0; i < S; i++ {
		fli[i] = make(pixels, S)
	}
	for i := 0; i < S; i++ {
		for j := 0; j < S; j++ {
			fli[i][j] = data[j][i]
		}
	}
	return fli
}

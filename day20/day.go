package day20

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/wistler/aoc-2020/internal"

	"github.com/wistler/aoc-2020/internal/io"
)

type pos struct {
	r, c int
}

// find transform needed to sync borders to target constraints
func findSyncState(src [4]int, target [4]int) (tileState, error) {
	for _, s := range states {
		sT := applyState(src, s)

		synced := true
		for i := 0; i < 4; i++ {
			if target[i] != -1 {
				if sT[i] != target[i] {
					synced = false
				}
			}
		}

		if synced {
			return s, nil
		}
	}
	return tileState{}, fmt.Errorf("Cannot sync: src: %v target: %v", src, target)
}

func applyState(src [4]int, state tileState) [4]int {
	var target [4]int
	for i := 0; i < 4; i++ {
		//  0 1 2 3  i
		// [0 1 2 3] rot=0
		// [1 2 3 0] rot=1
		// [2 3 0 1] rot=2
		// [3 0 1 2] rot=3
		r := (i + state.rotated) % 4

		//  0 1 2 3  i
		// [3 2 1 0] rot=0
		// [0 3 2 1] rot=1
		// [1 0 3 2] rot=2
		// [2 1 0 3] rot=3
		f := (state.rotated + 3 - i) % 4

		if state.flipped {
			target[i] = src[f]
		} else {
			target[i] = src[r]
		}
	}
	return target
}

func parseInput(input []string) map[int]tile {
	tiles := make(map[int]tile)
	var t tile
	t.id = 0
	t.rawData = make([][]bool, 10)
	tileLine := 0
	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			if t.id != 0 {
				tiles[t.id] = t
				t = tile{}
				t.rawData = make([][]bool, 10)
				tileLine = 0
			}
			continue
		}
		n, err := fmt.Sscanf(line, "Tile %d:", &t.id)
		if err == nil && n == 1 {
			continue
		}
		t.rawData[tileLine] = make([]bool, 10)
		for i, ch := range line {
			t.rawData[tileLine][i] = ch == '#'
		}
		tileLine++
	}
	return tiles
}

func stringsToBool(str string) [][]bool {
	data := [][]bool{}
	for _, line := range io.SplitOnNewLines(str) {
		lin := strings.TrimSpace(line)
		if lin == "" {
			continue
		}
		r := []bool{}
		for _, ch := range lin {
			r = append(r, ch == '#')
		}
		data = append(data, r)
	}
	return data
}

var monster [][]bool
var states []tileState = []tileState{
	{0, false},
	{1, false},
	{2, false},
	{3, false},
	{0, true},
	{1, true},
	{2, true},
	{3, true},
}

func init() {
	monster = stringsToBool(`
	..................#.
	#....##....##....###
	.#..#..#..#..#..#...
	`)
}

func jKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func zip(images ...[][]bool) [][]bool {
	z := [][]bool{}
	for k := 0; k < len(images[0]); k++ {
		zk := []bool{}
		for _, iD := range images {
			for _, iDk := range iD[k] {
				zk = append(zk, iDk)
			}
		}
		z = append(z, zk)
	}
	return z
}

func display(images ...[][]bool) {
	logPrefix := log.Prefix()
	log.SetPrefix("")
	for k := 0; k < len(images[0]); k++ {
		line := []string{}
		for _, iD := range images {
			line = append(line, toString(iD[k][:]))
		}
		log.Println(strings.Join(line, "  "))
	}

	log.SetPrefix(logPrefix)
}

func highlightDifference(left [][]bool, right [][]bool) {
	logPrefix := log.Prefix()
	log.SetPrefix("")
	for row := 0; row < len(left); row++ {
		line := make([]string, len(left[row]))
		for col := 0; col < len(left[row]); col++ {
			ch := "?"
			if left[row][col] == right[row][col] {
				ch = toStr(left[row][col])
			} else {
				ch = "\u001b[31mO\u001b[0m"
			}
			line = append(line, ch)
		}
		log.Println(strings.Join(line, ""))
	}
	log.SetPrefix(logPrefix)
}

func stitch(jigsaw map[pos]tile, tiles map[int]tile) [][]bool {
	image := [][]bool{}
	for i := 0; ; i++ {
		imageRow := [][][]bool{}
		for j := 0; ; j++ {
			jk := pos{i, j}
			t, ok := jigsaw[jk]
			if !ok {
				break
			}
			imageRow = append(imageRow, t.getData())
		}
		if len(imageRow) == 0 {
			break
		}
		for _, zr := range zip(imageRow...) {
			image = append(image, zr)
		}
	}
	return image
}

func findRoughness(image [][]bool, beast [][]bool) (roughness int, foundMonsters int, workImage [][]bool, err error) {
	workImage = make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		workImage[i] = make([]bool, len(image[i]))
		copy(workImage[i], image[i])
	}

	for i := 0; i < len(workImage)-len(beast)+1; i++ {
		for j := 0; j < len(workImage[0])-len(beast[0])+1; j++ {
			match := true

		Search:
			for k := 0; k < len(beast); k++ {
				for l := 0; l < len(beast[0]); l++ {
					if beast[k][l] {
						if !workImage[i+k][j+l] {
							match = false
							break Search
						}
					}
				}
			}
			if match {
				foundMonsters++
				for k := 0; k < len(beast); k++ {
					for l := 0; l < len(beast[0]); l++ {
						// erase the beast, so it doesn't miscount
						workImage[i+k][j+l] = workImage[i+k][j+l] && !beast[k][l]
					}
				}
			}
		}
	}
	if foundMonsters == 0 {
		err = errors.New("No monsters")
	} else {
		for i := 0; i < len(workImage); i++ {
			for j := 0; j < len(workImage[0]); j++ {
				if workImage[i][j] {
					roughness++
				}
			}
		}
	}
	return
}

func findRoughnessAndMosters(image [][]bool, beast [][]bool) (int, int) {
	r, m := 0, 0
	w := [][]bool{}
	var err error

	log.Print("Looking for monsters..\n\n")
	for _, st := range states {
		ns := rotate(image, st.rotated)
		if st.flipped {
			ns = flip(ns)
		}

		r, m, w, err = findRoughness(ns, beast)
		if err == nil {
			// display(ns, w)
			highlightDifference(ns, w)
			log.Println("Found", m, "monsters.")
			log.Println("Roughness:", r)
			break
		}
	}
	return r, m
}

func updatePairing(tiles map[int]tile) map[int]tile {
	for _, t1 := range tiles {
		pairedSides := 0
		t1b := t1.getBorders()

		for side := 0; side < 4; side++ {
			t1bs := t1b[side]

		L1:
			for _, t2 := range tiles {
				if t1.id == t2.id {
					continue
				}
				for _, s1 := range states {
					ts2 := t2.transform(s1)
					rb2 := ts2.getBorders()

					for bi := 0; bi < 4; bi++ {
						matched := internal.BoolsAreEqual(t1bs, rb2[bi])
						// if matched {
						// 	log.Println("Matching:", t1.id, toString(t1bs), toString(rb2[bi]), t2.id, side, si, bi, matched)
						// }
						if matched {
							pairedSides++
							t1.links[side] = t2.id
							break L1
						}
					}
				}
			}
		}
		tiles[t1.id] = t1
	}

	return tiles
}

func part1(input []string) int {
	log.SetPrefix("Day 20: Part 1: ")
	log.SetFlags(0)

	output := new(bytes.Buffer)
	log.SetOutput(output)

	tiles := parseInput(input)
	tiles = updatePairing(tiles)
	corners := []int{}
	for tid, t1 := range tiles {
		if t1.isCorner() {
			corners = append(corners, tid)
		}
	}
	log.Println("corners:", corners)

	println(output.String())
	output.Truncate(0)

	if len(corners) != 4 {
		panic("Incorrect corner count")
	}

	prod := 1
	for _, id := range corners {
		prod *= id
	}
	return prod
}

func part2(input []string) int {
	log.SetPrefix("Day 20: Part 2: ")
	log.SetFlags(0)

	output := new(bytes.Buffer)
	log.SetOutput(output)

	tiles := parseInput(input)
	tiles = updatePairing(tiles)
	corners := []int{}
	for tid, t1 := range tiles {
		if t1.isCorner() {
			corners = append(corners, tid)
		}
	}
	log.Println("corners:", corners)

	println(output.String())
	output.Truncate(0)

	if len(corners) != 4 {
		panic("Incorrect corner count")
	}

	// start at a corner (and position at top-left), and work our way across the image
	jigsaw := make(map[pos]tile)
	i, j := 0, 0
	I, J := 0, 0

	used := make([]int, len(tiles))
	cid := corners[0]

	log.Println("findingSyncState for tile:", cid)
	cs, err := findSyncState(tiles[cid].links, [4]int{0, -1, -1, 0})
	if err != nil {
		println(output.String())
		panic(err)
	}

	jigsaw[pos{i, j}] = tiles[cid].transform(cs)
	used = append(used, cid)
	log.Println("corner:", cid, "tile:", tiles[cid].id, "jKey:", jKey(i, j))
	println(output.String())

	j++
	for {
		constrain := [4]int{-1, -1, -1, -1}

		jid := -1

		jkw := pos{i, j - 1}
		if ckw, ok := jigsaw[jkw]; ok {
			constrain[3] = ckw.id

			// east-facing constraint on west tile
			syncedPairing := ckw.links
			jid = syncedPairing[1]
		}

		jkn := pos{i - 1, j}
		if ckn, ok := jigsaw[jkn]; ok {
			constrain[0] = ckn.id

			// south-facing constraint on north tile
			syncedPairing := ckn.links
			jidn := syncedPairing[2]

			if jid != -1 && jidn != jid {
				log.Println(jigsaw)
				log.Println("i:", i, "j:", j, "jid:", jid, "jidn:", jidn)

				println(output.String())
				panic("Combination error")
			}
			jid = jidn
		}

		if i == 0 {
			constrain[0] = 0 // north border
		}
		if j == 0 {
			constrain[3] = 0 // west border
		}

		if jid == -1 {
			log.Println("jigsaw:", jigsaw)
			log.Println("i:", i, "j:", j, "jid:", jid)

			println(output.String())
			panic("Error finding next piece")
		}

		if ok, _ := internal.ContainsNumber(used, jid); ok {
			log.Println("jigsaw:", jigsaw)
			log.Println("i:", i, "j:", j, "jid:", jid)

			println(output.String())
			panic("Error finding next piece")
		}

		jp := tiles[jid].links
		jstate, err := findSyncState(jp, constrain)
		if err != nil {
			log.Println("jigsaw:", jigsaw)
			log.Println("i:", i, "j:", j, "jid:", jid)

			println(output.String())
			panic(err)
		}
		jigsaw[pos{i, j}] = tiles[jid].transform(jstate)
		used = append(used, jid)
		jb := applyState(jp, jstate)
		if jb[1] == 0 { // east border
			log.Println("Found east border @", jkw)

			i++
			if J == 0 {
				J = j + 1
			}
			j = 0
			if jb[2] == 0 {
				// south-east corner of jigsaw, the end
				I = i
				log.Println("Found south-east border @", jkw)
				break
			}
		} else {
			j++
		}
		output.Truncate(0)
	}

	log.Println("jigsaw size:", I, J)
	for i := 0; i < I; i++ {
		images := make([][][]bool, J)
		for j := 0; j < J; j++ {
			images[j] = jigsaw[pos{i, j}].rawData
		}
		display(images...)
		output.WriteString("\n")
	}

	fullImage := stitch(jigsaw, tiles)
	roughness, _ := findRoughnessAndMosters(fullImage, monster)
	println(output.String())
	output.Truncate(0)

	log.Printf("Answer: %v", roughness)
	println(output.String())
	output.Truncate(0)
	return roughness
}

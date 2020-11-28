//
// Sample app to test using a package from same workspace
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wistler/aoc-2020/lib"
)

// Sum returns sum of vectors
func Sum(pts []lib.Vec) lib.Vec {
	sum := lib.MakeVector(0, 0)
	for _, pt := range pts {
		sum = sum.Add(pt)
	}
	return sum
}

func main() {
	log.SetPrefix("Sample: ")
	log.SetFlags(0)

	args := os.Args[1:]
	// log.Printf("Args: %v\n", args)
	if len(args) == 0 {
		log.Fatal("Need args. Usage: sample [1,2] [3,4]")
	}

	pts := make([]lib.Vec, len(args))

	for i, arg := range args {
		var x, y int
		n, err := fmt.Sscanf(arg, "[%d,%d]", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatalf("Ill-formed Argument: %q\n", arg)
		}

		pts[i] = lib.MakeVector(x, y)
	}

	log.Printf("Vectors = %v\n", pts)
	log.Printf("Sum of vectors = %v\n", Sum(pts))
}

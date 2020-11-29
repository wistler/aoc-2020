//
// Sample app to test using a package from same workspace
//

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wistler/aoc-2020/internal/vector"
)

// Sum returns sum of vectors
func Sum(pts []vector.Vec) vector.Vec {
	sum := vector.Make(0, 0)
	for _, pt := range pts {
		sum.Add(pt)
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

	pts := make([]vector.Vec, len(args))

	for i, arg := range args {
		var x, y float64
		n, err := fmt.Sscanf(arg, "[%f,%f]", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatalf("Ill-formed Argument: %q\n", arg)
		}

		pts[i] = vector.Make(x, y)
	}

	log.Printf("Vectors = %v\n", pts)
	log.Printf("Sum of vectors = %v\n", Sum(pts))
}

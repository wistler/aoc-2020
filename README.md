# README

Learning repo for puzzles of `Advent Of Code 2020`, using Go lang.

To run the solution of a given day, from the project root dir, run:

```sh
> go test -v ./dayNN
```

## Why `go test` ?

I'm still figuring out how to use `go run` meaningfully. I'd found that `go run` doesn't work nicely when importing packages from relative paths, and that we'd have use `go build` to pick up the package dependency.

Alternatively, I'd like to try running the solutions as `go "scripts"`. Still reading up on that.

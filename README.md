# README

![Tests](https://github.com/wistler/aoc-2020/workflows/Tests/badge.svg)

Learning repo for puzzles of `Advent Of Code 2020`, using Go lang.

To run the solution of a given day, from the project root dir, run:

```sh
> go test -v ./dayNN

# To run solutions for all days, use wildcard format:
> go test -v ./day*

# To run tests with sample data for all days, use:
> go test -v ./day* -run Sample

# To run tests with real data for all days, use:
> go test -v ./day* -run Real
```

## Implementation choices
- Choosing to panic, in favor of script-like coding. These are short snippets that don't benefit much from error-handling.

## Why `go test` ?
I'm still figuring out how to use `go run` meaningfully.
- Relative path to input file must be from project root dir. This limitation doesn't exist for `go test`.
- Assertions with sample data should be "tests". So, I have to use `go test` anyway.

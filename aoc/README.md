# AoC

This package provides some helper functions to make the Advent of Code solutions easier, faster and shorter to solve.

## Examples

```go
package main

import (
    "github.com/kkrypt0nn/swiss-army-knife/aoc"
)

func main() {
    str := "hello world"
    substr := "world"
    fmt.Println(aoc.ContainsExactly(str, substr))
    // Output: false
    str = "world"
    substr = "dlrow"
    fmt.Println(aoc.ContainsExactly(str, substr))
    // Output: true
}
```

```go
package main

import (
    "github.com/kkrypt0nn/swiss-army-knife/aoc"
)

func main() {
    myMap := map[int]int{
        0: 5,
        1: 10,
        2: 15,
    }
    fmt.Println(aoc.MapSum(myMap))
    // Output: 30
}
```

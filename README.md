# Axway Go workshop
This repository contains guidelines for Axway's go workshop.

# Introduction
This document will quickly cover (reference) some of the information in (gotour)[https://tour.golang.org]

```go
package main

import (
  "fmt"
)

const answer = 42

func doSwap(x, y int64) (int64, int64) {
  return y, x
}

// Swap swaps the two provided values
func Swap(x, y int64) (int64, int64) {
  return doSwap(x, y)
}

func main() {
  v1 := 1 // short initialization. Equal to var v1 = 1
  var v2 int64 // declaration of variable with type. The value is the default value for the given type. In this case 0
  v2 = 2
  var y2, y1 = Swap(v1, v2)
  fmt.Printf("%d, %d", y1, y2)
}
```

### Types
```
bool (false)

string ("")

int  int8  int16  int32  int64 (0)
uint uint8 uint16 uint32 uint64 uintptr (0)

byte // alias for uint8 (0)

rune // alias for int32 (0)
     // represents a Unicode code point

float32 float64 (0)

complex64 complex128 (0)
```

#### Conversion
Conversion is explicit. Can't just assign values of different types to eachother (e.g. `int` to `float`)
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

#### Inference
```go
// from constant
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
s := "sometext"   // string

// from other variable
var i int
j := i // j is an int, because i is an int
```

# Flow control
## Loops
Go only has `for` structure for looping over constructs. The typical format follows the one of other C style languages:
```go
for var i = 0; i < 10; i ++ {
  doStuff(i)
}
```

Note that the first (init) and last (post) statements in the for loop declaration are optional and can be omitted.
```go
sum := 1
for ; sum < 100; {
  sum += sum
}
```

This is more clearly written as:
```go
sum := 1
for sum < 100 {
  sum += sum
}
```
And for an infinite loop you can just omit the condition part:
```go
for {
  doStuff()
}
```
## Conditional

### if
Basic construct:
```go
if x < 0 {
  dealWithIt()
} else {
  dont()
}
```

`if` with initialization statement:

```go
if x := 0; x < 0 {
  yesStupidExample()
}
```

### switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression. 

```go
fmt.Print("Go runs on ")
switch os := runtime.GOOS; os {
case "darwin":
  fmt.Println("OS X.")
case "linux":
  fmt.Println("Linux.")
default:
  // freebsd, openbsd,
  // plan9, windows...
  fmt.Printf("%s.", os)
}
```

Note that unlike other languages switch cases doesn't fall through. You need an explicit fallthrough statement if you want that behaviour
```go
switch { // = switch true - always looks through the case statements
case "a":
  fmt.Println("a")
  fallthrough // causes it to continue evaluating the other cases
case "b":
  fmt.Println("b")
}
```

## Defer
`defer` allows a function call to be delayed until the end of the current function scope. It's typically used to delay closing closable objects close to the place they were opened/initialized.

```go
func func1() {
  f := createFile("filename.txt")
  defer f.Close() // this will be called when ReadStuff finishes
                  // and we are about to exit the scope of func1
  f.ReadStuff()
}
```

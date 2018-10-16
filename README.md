# Axway Go workshop
This repository contains guidelines for Axway's go workshop.

# Introduction
TODO

# Go Tour
https://tour.golang.org

## Structure
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand"
package comprises files that begin with the statement package rand. 
## Types
### Basic types
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
#### Zero values
```go
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
an empty struct with corresponding zero values for its fields
nil for pointers, slices and maps
```
#### Type conversion
```go
var i int = 42 // int can be omitted as it's inferred
var f float32 = float(32) // explicit conversion    
```

### More types
#### Structs
```go
// Point represents a point in a two dimensional space
type Point struct {
    X int
    Y int
}
var p := Point{1, 1} // Could be written verbosely as := Pointer{X: 1, Y: 1}
fmt.Println(p.X) // access field
p.Y = 2 // assign to field
fmt.Println(p.Y)
```
#### Pointers
```go
var p *int     // declare p as pointer to an integer
i := 42
p = &i         // get the address of i and assign it to p
fmt.Printf(*p) // dereference the pointer and get the value stored at its address

var p := &Point{1, 1} // pointer to a Point instance
```
#### Arrays
```go
// fixed size. Can't be resized
[3]bool{true, true, false}
```  
#### Slices
```go
// a view over arrays. Can be resized (internally a new array is allocated and the values - copied)
[]bool{true, true, false} // creates an array and gets a slice reference to it
// Creating a slice:
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
c := []int{1, 2, 3, 4, 5} // len(c)=5, cap(c)=5
// appending to a slice:
a = append(a, 2, 3, 4)
// slicing
x := a[1:3] // 1 and/or 3 can be omitted and they are automatically set to 0/len
```  
#### Maps
```go
// create
var m = make(map[string]Point)
m["point1"] = Point{1, 1}
m["point2"] = Point{2, 2}

// OR as literal
var m = map[string]Point {
	"point1": Point{1, 1},
	"point2": Point{2, 2},
}

// access
fmt.Println(m["point1"])

// update
m["point2"] = Point{3, 3}

// delete
delete(m, "point2")

// check if key exists in map:
v, ok = m["point2"] // ok is boolean value that will be true if the key exists and false otherwise.
                    // If it exists v has the value for this key
```

#### Functions (yes, it's a type)
```go
// compute accepts a function (that receives two integers and returns an integer) and returns an integer
func compute(fn func(int, int) int) int { // :troll_face:
	return fn(3, 4)
}

// power accepts two integers and returns an integer (thus matching the signature that compute expects)
func power (x, y int) int {
	return math.Pow(x, y)
}

fmt.Println(compute(power))
```

## Basics
### File structure
###  Variables
```go
var x int
x = 9
var y = 9
z := 9
const pi = 3.14159265359
```
### Functions
```go
func Add (x, y int) int {
	return x + y
}

// can also return more than one value

func splitInTwo(s string, sep string) (string, string) {
	result := strings.SplitN(s, sep, 2)
	return result[0], result[1]
}
```
### Defer
   ```go
   func ParseFile(filename string) string {
   	f := ioutil.ReadAll(filename)
   	defer f.Close() // this will be called after the return
   	
   	// ... do stuff with the file content
   	
   	return "success" // yeah. ok!
   }
   ```
### Exports
```go
// Power is exported (part of the module's API)
func PowerTwo (x int) int {
	return power(x, 2)
}

// power is not exported (i.e. invisible to outside packages)
func power (x, y int) int {
	return math.Pow(x, y)
}
```

## Flow control
### For (one word to loop them all)
```go
for [pre]; [condition]; [post] { // all optional
	//...
}
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
for i < 10 { // basically a while
	fmt.Println(i++)
}
for { // while true
	// boom
}
```
#### Range
```go
// for slices and maps
slice := []int{1, 2, 3}
for index, elem := range slice {
    fmt.Printf("%d: %d", index, elem)
}

m := map[string]string{"k1": "v1", "k2": "v2"}
for k, v := range m {
	fmt.Printf("%s: %s", k, v)
}
```
### If
```go
if [pre]; condition {
	//...
} else if ... {
	//...
} else {
	
}
if x := y; x < 10 { // x is only available in 'if's scope
	// ...
}
```
### Switch (easier if-else chains. No automatic fallthrough!)
```go
switch [pre]; [condition] {
case A:
	doA()
case B:
	doB()
    fallthrough // explicitly do C as well
case C:
	doC()
default:
	doDefault()
}

// e.g.  pre + switch on os
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

// e.g. switch without condition
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

## Methods
```go
// Value receiver
type Point struct {
	X, Y int
}

func (p Point) Pretty() string { // Cannot modify p
	return "X: " + p.X + ", Y: " + p.Y
}

v := Point{1, 1}
fmt.Println(v.Pretty())

// Pointer receiver
type Point struct {
	X, Y int
}

func (p *Point) MoveUp() { // Modifies p
	p.X + 1
}
func (p *Point) MoveDown() {
 	p.X - 1
}

v := Point{1, 1}
v.MoveUp()
fmt.Println(v.X)
```

## Interfaces
### Definition
An interface type is defined as a set of method signatures.
```go
// Logger is an interface that has a single method, accepting a string
type Logger interface {
	Info(string)
}

// MyLogger implements the Logger interface
type MyLogger struct {
	Prefix string
}

// Info has the exact signature Logger requires so that it's implemented
func (l MyLogger) Info(msg string) {
	fmt.Println(l.Prefix + ": " + msg)
}

// YourLogger does NOT implement the Logger interface (no Info method)
type YourLogger struct {
	Prefix string
}

// logMiddleware accepts anything that implements the Logger interface and uses it to log a message
func logMiddleware(logger Logger) {
	logger.Info("Hi")
}

myLogger := MyLogger{"InfoLvl"}
yourLogger := YourLogger{"InfoLvl"}

logMiddleware(myLogger)   // will print "InfoLvl: Hi"
logMiddleware(yourLogger) // will panic because yourLogger doesn't implement Logger (no Info method)
``` 

### Type assertion/casting
```go
myLog, ok := myLogger.(Logger) // ok is true, because myLogger implements Logger.
                               // myLog is casted to Logger
yourLog, ok := yourLogger.(Logger) // ok is false, because yourLogger does not implements Logger.
                                   // yourLog is assigned its default value (empty struct)
yourLog2 := yourLogger.(Logger) // panics
```

## Errors
```go
i, err := strconv.Atoi("42") // tries to convert the string to an integer
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer: ", i)
```

## Goroutines
```go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") // "go" starts a new goroutine
	say("hello")
}
```
## Channels
Declaration
```go
ch := make(chan int)
ch <- x    // Send v to channel ch.
y := <-ch  // Receive from ch, and
           // assign value to v.
```

For example:
```go
// returnAfter waits for sec seconds and sends a value over the channel c
func returnAfter (sec int, c chan bool) {
	time.Sleep(time.Duration(sec) * time.Second) // notice that sec (type int) must explicitly be converted
	                                             // to type time.Duration
	c <- true // sends a value true over the channel
}
func main() {
	c := make(chan bool)
	go returnAfter(2, c)
	val := <-c // blocks until there's something to receive from the channenl
	fmt.Println("done: " + val) // done: true
}
```

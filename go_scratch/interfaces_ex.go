package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type BINARY uint

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Internally following happens:
// g := geometry(r) // depending on the type of r, g will be of type rect or circle
// g hold two values: type (rect or circle) and value of the passed object
// g.area() // What this does is, it checks the type of g and calls the area() method of that type and passes the pointer to value of g to that method and dereferences it Internally
func measure(g geometry) {
	fmt.Printf("%v, %T \n", g, g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

// func main() {
// 	r := rect{width: 3, height: 4}
// 	c := circle{radius: 5}
//
// 	measure(r)
// 	measure(c)
//
// 	detectCircle(r)
// 	detectCircle(c)
//
// 	geo := geometry(r)
// 	fmt.Printf("%v, %T\n", geo, geo)
//
// 	// ReadAndClose(r, []byte("test"))
//
// 	co := container{base: base{num: 1}, str: "test"}
// 	fmt.Println(co.num, co.str)
// 	s := []int{1, 2, 3, 4, 5}
//
// 	// Create a new slice with a copy of the elements from s
// 	copiedSlice := append(s[:0:0], s...)
//
// 	// Modify the original slice (to show that it doesn't affect copiedSlice)
// 	s[0] = 100
//
// 	// Print both slices
// 	fmt.Println("Original slice:", s)         // Output: [100 2 3 4 5]
// 	fmt.Println("Copied slice:", copiedSlice) // Output: [1 2 3 4 5]
// }

type ReadCloser interface {
	Read(b []byte) (n int, err os.LinkError)
	Close()
}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err os.LinkError) {
	for len(buf) > 0 {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateActive
	StateTerminated
)

type base struct {
	num int
}
type container struct {
	base
	str string
}

func InstantiatedClone1(s []string) []string {
	return append(s[:0:0], s...)
}
func Clone1[E any](s []E) []E {
	return append(s[:0:0], s...)
}

// MySlice is a slice of strings with a special String method.
type MySlice []string

// String returns the printable version of a MySlice value.
func (s MySlice) String() string {
	return strings.Join(s, "+")
}
func PrintSorted(ms MySlice) string {
	c := Clone3(ms)
	slices.Sort(c)
	return c.String() // FAILS TO COMPILE
}

// Type parameters --> S ~[]E, E any
// Type arguments --> S
func Clone3[S ~[]E, E any](s S) S {
	return append(s[:0:0], s...)
}

func (s MySlice) iterator() func(func(string) bool) {
	return func(yield func(string) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

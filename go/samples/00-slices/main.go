package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

type sliceHeader struct {
	length   int
	zerothEl *byte
}

func printSep() {
	fmt.Println("==========")
}

func addOneToEachEl(slice []int) {
	for i := range slice {
		slice[i]++
	}
}

func addOneToEachElAndReturn(slice []int) []int {
	for i := range slice {
		slice[i]++
	}
	return slice
}

// slice pointer
type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) TruncateAtFinalSlashUseVar() path {
	i := bytes.LastIndex(p, []byte("/"))
	if i >= 0 {
		p = p[0:i]
	}
	return p
}

func (p path) ToUpperVal() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	var buffer [256]byte
	slice := sliceHeader{
		length:   5,
		zerothEl: &buffer[50],
	}

	for index := range slice.length {
		uindex := uint(index)
		curr := unsafe.Add(unsafe.Pointer(slice.zerothEl), uindex)
		currVal := *(*int)(curr)
		fmt.Printf("start=%p, index=%v, pointer=%p, pointed value=%v\n", slice.zerothEl, index, curr, currVal)
	}
	printSep()

	ts := []string{"dat", "123", "456", "789", "dummy test one"}
	ss1 := ts[1:5]
	ss2 := ss1[1 : len(ss1)-1]
	fmt.Printf("ss1=[%v], ss2=[%v]\n", ss1, ss2)

	iArray1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	iSlice1 := iArray1[2:6]
	fmt.Printf("before iArray1=%v, iSlice1=%v\n", iArray1, iSlice1)
	addOneToEachEl(iSlice1)
	fmt.Printf("after iArray1=%v, iSlice1=%v\n", iArray1, iSlice1)

	iArray2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	iSlice2 := iArray2[2:6]
	fmt.Printf("before iArray2=%v, iSlice2=%v\n", iArray2, iSlice2)
	addOneToEachElAndReturn(iSlice2)
	fmt.Printf("after iArray2=%v, iSlice2=%v\n", iArray2, iSlice2)

	printSep()
	p1 := path("/usr/local/bin")
	fmt.Printf("before p1=%s\n", p1)
	p1.TruncateAtFinalSlash()
	fmt.Printf("after p1=%s\n", p1)

	p2 := path("/usr/local/help")
	fmt.Printf("before p2=%s\n", p2)
	p2Ret := p2.TruncateAtFinalSlashUseVar()
	fmt.Printf("after p2=%s, p2Ret=%s\n", p2, p2Ret)

	printSep()
	p3 := path("/usr/local/bin")
	fmt.Printf("before method=ToUpperVal, p3=%s\n", p3)
	p3.ToUpperVal()
	fmt.Printf("after method=ToUpperVal, p3=%s\n", p3)
}

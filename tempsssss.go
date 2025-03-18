package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type BIT2D struct {
	n      int
	coords [][]int 
	tree   [][]int 
}

func NewBIT2D(n int) *BIT2D {
	b := &BIT2D{
		n:      n,
		coords: make([][]int, n+1),
		tree:   make([][]int, n+1),
	}
	return b
}

func (b *BIT2D) addCoord(pos, val int) {
	b.coords[pos] = append(b.coords[pos], val)
}

func (b *BIT2D) init() {
	for i := 1; i <= b.n; i++ {
		if len(b.coords[i]) > 0 {
			sort.Ints(b.coords[i])

			unique := b.coords[i][:0]
			for _, x := range b.coords[i] {
				if len(unique) == 0 || unique[len(unique)-1] != x {
					unique = append(unique, x)
				}
			}
			b.coords[i] = unique
		}
		b.tree[i] = make([]int, len(b.coords[i])+1)
	}
}

func (b *BIT2D) update(pos, val, delta int) {
	for i := pos; i <= b.n; i += i & -i {

		j := sort.SearchInts(b.coords[i], val)

		if j < len(b.coords[i]) && b.coords[i][j] == val {

			for k := j + 1; k < len(b.tree[i]); k += k & -k {
				b.tree[i][k] += delta
			}
		}
	}
}

func (b *BIT2D) query(pos, val int) int {
	res := 0
	for i := pos; i > 0; i -= i & -i {

		j := sort.SearchInts(b.coords[i], val+1)

		for k := j; k > 0; k -= k & -k {
			res += b.tree[i][k]
		}
	}
	return res
}

func (b *BIT2D) QueryRange(l, r, val int) int {
	return b.query(r, val) - b.query(l-1, val)
}

type Query struct {
	typ int
	x   int  
	c   byte 
	d   int  
	k   int  
}

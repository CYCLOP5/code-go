type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+1),
		n:    n,
	}
}

func (b *BIT) update(i, val int) {
	for ; i <= b.n; i += i & -i {
		b.tree[i] += val
	}
}

func (b *BIT) query(i int) int {
	sum := 0
	for ; i > 0; i -= i & -i {
		sum += b.tree[i]
	}
	return sum
}

func cmpress(x, offset int) int {
	return x + offset + 1
}
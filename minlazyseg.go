
func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
func abs(x int) int    { if x < 0 { return -x }; return x }
type node struct {
	minval int
	lazy   int 
}

type SegTree struct {
	tree []node
	size int
}

func NewSegTree(size int) *SegTree {
	st := &SegTree{}
	st.size = size
	st.tree = make([]node, 4*size)
	st.build(1, 0, size-1)
	return st
}

func (st *SegTree) build(v, tl, tr int) {
	st.tree[v] = node{minval: posinf, lazy: posinf}
	if tl == tr {
		return
	}
	tm := (tl + tr) / 2
	st.build(2*v, tl, tm)
	st.build(2*v+1, tm+1, tr)
}

func (st *SegTree) push(v int) {
	if st.tree[v].lazy == posinf {
		return
	}
	val := st.tree[v].lazy
	for _, c := range []int{2 * v, 2*v + 1} {
		if st.tree[c].minval > val {
			st.tree[c].minval = val
		}
		if st.tree[c].lazy > val {
			st.tree[c].lazy = val
		}
	}
	st.tree[v].lazy = posinf
}

func (st *SegTree) updrange(v, tl, tr, l, r, value int) {
	if l > r {
		return
	}
	if l == tl && tr == r {
		st.tree[v].minval = min(st.tree[v].minval, value)
		st.tree[v].lazy = min(st.tree[v].lazy, value)
		return
	}
	st.push(v)
	tm := (tl + tr) / 2
	st.updrange(2*v, tl, tm, l, min(r, tm), value)
	st.updrange(2*v+1, tm+1, tr, max(l, tm+1), r, value)
	st.tree[v].minval = min(st.tree[2*v].minval, st.tree[2*v+1].minval)
}

func (st *SegTree) rangequery(v, tl, tr, l, r int) int {
	if l > r {
		return posinf
	}
	if l <= tl && tr <= r {
		return st.tree[v].minval
	}
	st.push(v)
	tm := (tl + tr) / 2
	res_l := st.rangequery(2*v, tl, tm, l, min(r, tm))
	res_r := st.rangequery(2*v+1, tm+1, tr, max(l, tm+1), r)
	return min(res_l, res_r)
}

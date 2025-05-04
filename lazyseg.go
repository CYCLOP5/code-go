type node struct {
	maxval  int
	heurval int
}
type SegTree struct {
	tree []node
	size int
}

func NewSegTree(size int) *SegTree {
	st := &SegTree{}
	st.size = size
	st.tree = make([]node, 4*size)
	st.build(1, 1, size)
	return st
}
func (st *SegTree) build(v, tl, tr int) {
	st.tree[v] = node{maxval: neginf, heurval: neginf}
	if tl == tr {
		return
	}
	tm := (tl + tr) / 2
	st.build(2*v, tl, tm)
	st.build(2*v+1, tm+1, tr)
}
func (st *SegTree) push(v, tl, tr int) {
	if st.tree[v].heurval == neginf || tl == tr {
		return
	}
	lazyVal := st.tree[v].heurval
	st.tree[2*v].maxval = max(st.tree[2*v].maxval, lazyVal)
	st.tree[2*v+1].maxval = max(st.tree[2*v+1].maxval, lazyVal)
	st.tree[2*v].heurval = max(st.tree[2*v].heurval, lazyVal)
	st.tree[2*v+1].heurval = max(st.tree[2*v+1].heurval, lazyVal)
	st.tree[v].heurval = neginf
}
func (st *SegTree) updrange(v, tl, tr, l, r, value int) {
	if l > r || tr < l || tl > r {
		return
	}
	if l <= tl && tr <= r {
		st.tree[v].maxval = max(st.tree[v].maxval, value)
		st.tree[v].heurval = max(st.tree[v].heurval, value)
		return
	}
	st.push(v, tl, tr)
	tm := (tl + tr) / 2
	st.updrange(2*v, tl, tm, l, r, value)
	st.updrange(2*v+1, tm+1, tr, l, r, value)
	st.tree[v].maxval = max(st.tree[2*v].maxval, st.tree[2*v+1].maxval)
}
func (st *SegTree) pointerquery(v, tl, tr, pos int) int {
	if tl == tr {
		return st.tree[v].maxval
	}
	st.push(v, tl, tr)
	tm := (tl + tr) / 2
	var res int
	if pos <= tm {
		res = st.pointerquery(2*v, tl, tm, pos)
	} else {
		res = st.pointerquery(2*v+1, tm+1, tr, pos)
	}
	return res
}
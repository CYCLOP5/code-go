type segtreenode struct {
	maxavail int
}
type segtree struct {
	nodes []segtreenode
	n     int
}

func newseg(size int) *segtree {
	st := &segtree{
		nodes: make([]segtreenode, 4*(size+1)),
		n:     size,
	}
	st.build(1, 1, size)
	return st
}
func (st *segtree) build(nidx, l, r int) {
	if l == r {
		st.nodes[nidx] = segtreenode{maxavail: l}
		return
	}
	mid := (l + r) / 2
	st.build(2*nidx, l, mid)
	st.build(2*nidx+1, mid+1, r)
	st.nodes[nidx].maxavail = maxint(st.nodes[2*nidx].maxavail, st.nodes[2*nidx+1].maxavail)
}
func (st *segtree) Update(taridx int, avail bool) {
	if taridx < 1 || taridx > st.n {
		return
	}
	st.updrecur(1, 1, st.n, taridx, avail)
}
func (st *segtree) updrecur(nidx, l, r, taridx int, avail bool) {
	if l == r {
		if avail {
			st.nodes[nidx].maxavail = l
		} else {
			st.nodes[nidx].maxavail = 0
		}
		return
	}
	mid := (l + r) / 2
	if taridx <= mid {
		st.updrecur(2*nidx, l, mid, taridx, avail)
	} else {
		st.updrecur(2*nidx+1, mid+1, r, taridx, avail)
	}
	st.nodes[nidx].maxavail = maxint(st.nodes[2*nidx].maxavail, st.nodes[2*nidx+1].maxavail)
}
func (st *segtree) querymax(queryL, queryR int) int {
	queryL = maxint(1, queryL)
	queryR = minint(st.n, queryR)
	if queryL > queryR {
		return 0
	}
	return st.querymaxrecur(1, 1, st.n, queryL, queryR)
}
func (st *segtree) querymaxrecur(nidx, l, r, queryL, queryR int) int {
	if r < queryL || l > queryR {
		return 0
	}
	if st.nodes[nidx].maxavail == 0 {
		return 0
	}
	if queryL <= l && r <= queryR {
		return st.nodes[nidx].maxavail
	}
	mid := (l + r) / 2
	rightMax := st.querymaxrecur(2*nidx+1, mid+1, r, queryL, queryR)
	leftMax := st.querymaxrecur(2*nidx, l, mid, queryL, queryR)
	return maxint(leftMax, rightMax)
}
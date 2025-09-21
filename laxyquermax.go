
type node struct {
	maxv    int
	ladd int
	minp    int
}

type segtree struct {
	tree []node
	n    int
}

func newtree(size int) *segtree {
	st := &segtree{
		tree: make([]node, 4*size),
		n:    size,
	}
	st.build(1, 0, st.n-1)
	return st
}

func (st *segtree) build(v, tl, tr int) {
	st.tree[v] = node{maxv: 0, ladd: 0, minp: 0}
	if tl != tr {
		tm := (tl + tr) / 2
		st.build(2*v, tl, tm)
		st.build(2*v+1, tm+1, tr)
	}
}

func (st *segtree) push(v int) {
	if st.tree[v].ladd != 0 {
		st.tree[2*v].maxv += st.tree[v].ladd
		st.tree[2*v].ladd += st.tree[v].ladd
		st.tree[2*v+1].maxv += st.tree[v].ladd
		st.tree[2*v+1].ladd += st.tree[v].ladd
		st.tree[v].ladd = 0
	}
}

func (st *segtree) upd(v, tl, tr, l, r, add int) {
	if l > r {
		return
	}
	if l == tl && r == tr {
		st.tree[v].maxv += add
		st.tree[v].ladd += add
	} else {
		st.push(v)
		tm := (tl + tr) / 2
		st.upd(2*v, tl, tm, l, min(r, tm), add)
		st.upd(2*v+1, tm+1, tr, max(l, tm+1), r, add)
		st.tree[v].maxv = max(st.tree[2*v].maxv, st.tree[2*v+1].maxv)
	}
}

func (st *segtree) set(v, tl, tr, pos, val int) {
	if tl == tr {
		st.tree[v].maxv = val
		st.tree[v].ladd = 0
	} else {
		st.push(v)
		tm := (tl + tr) / 2
		if pos <= tm {
			st.set(2*v, tl, tm, pos, val)
		} else {
			st.set(2*v+1, tm+1, tr, pos, val)
		}
		st.tree[v].maxv = max(st.tree[2*v].maxv, st.tree[2*v+1].maxv)
	}
}

func (st *segtree) updp(v, tl, tr, pos, p_val int) {
	if tl == tr {
		st.tree[v].minp = p_val
	} else {
		tm := (tl + tr) / 2
		if pos <= tm {
			st.updp(2*v, tl, tm, pos, p_val)
		} else {
			st.updp(2*v+1, tm+1, tr, pos, p_val)
		}
		st.tree[v].minp = min(st.tree[2*v].minp, st.tree[2*v+1].minp)
	}
}

func (st *segtree) querymax(v, tl, tr, minprfix int) int {
	if st.tree[v].minp >= minprfix {
		return -inf
	}
	if tl == tr {
		return st.tree[v].maxv
	}

	st.push(v)
	tm := (tl + tr) / 2

	rezl := st.querymax(2*v, tl, tm, minprfix)

	minpfix := min(minprfix, st.tree[2*v].minp)
	rezr := st.querymax(2*v+1, tm+1, tr, minpfix)

	return max(rezl, rezr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
  if a < b { 
    return a 
  }
  return b 
}


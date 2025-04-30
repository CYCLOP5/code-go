type SegTreeNode struct {
	maxAvailable int 
}
type SegmentTree struct {
	nodes []SegTreeNode
	n     int 
}
func NewSegmentTree(size int) *SegmentTree {
	st := &SegmentTree{
		nodes: make([]SegTreeNode, 4*(size+1)),
		n:     size,
	}
	st.build(1, 1, size) 
	return st
}
func (st *SegmentTree) build(nodeIdx, L, R int) {
	if L == R {
		st.nodes[nodeIdx] = SegTreeNode{maxAvailable: L}
		return
	}
	mid := (L + R) / 2
	st.build(2*nodeIdx, L, mid)
	st.build(2*nodeIdx+1, mid+1, R)
	st.nodes[nodeIdx].maxAvailable = maxInt(st.nodes[2*nodeIdx].maxAvailable, st.nodes[2*nodeIdx+1].maxAvailable)
}
func (st *SegmentTree) Update(targetIdx int, available bool) {
	if targetIdx < 1 || targetIdx > st.n {
		return 
	}
	st.updateRecursive(1, 1, st.n, targetIdx, available)
}
func (st *SegmentTree) updateRecursive(nodeIdx, L, R, targetIdx int, available bool) {
	if L == R {
		if available {
			st.nodes[nodeIdx].maxAvailable = L 
		} else {
			st.nodes[nodeIdx].maxAvailable = 0 
		}
		return
	}
	mid := (L + R) / 2
	if targetIdx <= mid {
		st.updateRecursive(2*nodeIdx, L, mid, targetIdx, available)
	} else {
		st.updateRecursive(2*nodeIdx+1, mid+1, R, targetIdx, available)
	}
	st.nodes[nodeIdx].maxAvailable = maxInt(st.nodes[2*nodeIdx].maxAvailable, st.nodes[2*nodeIdx+1].maxAvailable)
}
func (st *SegmentTree) QueryMaxAvailable(queryL, queryR int) int {
	queryL = maxInt(1, queryL)
	queryR = minInt(st.n, queryR)
	if queryL > queryR { 
		return 0
	}
	return st.queryMaxAvailableRecursive(1, 1, st.n, queryL, queryR)
}
func (st *SegmentTree) queryMaxAvailableRecursive(nodeIdx, L, R, queryL, queryR int) int {
	if R < queryL || L > queryR {
		return 0 
	}
	if st.nodes[nodeIdx].maxAvailable == 0 {
		return 0
	}
	if queryL <= L && R <= queryR {
		return st.nodes[nodeIdx].maxAvailable 
	}
	mid := (L + R) / 2
	rightMax := st.queryMaxAvailableRecursive(2*nodeIdx+1, mid+1, R, queryL, queryR)
	leftMax := st.queryMaxAvailableRecursive(2*nodeIdx, L, mid, queryL, queryR)
	return maxInt(leftMax, rightMax)
}
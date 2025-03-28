var tree []Node

func merge(left, right Node) Node {
	res := Node{make([]uint8, K)}
	for i := 0; i < K; i++ {
		if left.me[i] > right.me[i] {
			res.me[i] = left.me[i]
		} else {
			res.me[i] = right.me[i]
		}
	}
	return res
}

func build(idx, L, R int, initarr []int) {
	tree[idx] = Node{make([]uint8, K)}
	if L == R {
		val := initarr[L-1]
		if val > 0 && val <= ipmax {
			factors := pfacs[val]
			for _, f := range factors {
				if pIdx, ok := ptindex[f.p]; ok {
					tree[idx].me[pIdx] = uint8(f.e)
				}
			}
		}
		return
	}
	M := L + (R-L)/2
	build(2*idx, L, M, initarr)
	build(2*idx+1, M+1, R, initarr)
	tree[idx] = merge(tree[2*idx], tree[2*idx+1])
}

func update(idx, L, R, targetIdx, newValue int) {
	if L == R {
		tree[idx].me = make([]uint8, K)
		if newValue > 0 && newValue <= ipmax {
			factors := pfacs[newValue]
			for _, f := range factors {
				if pIdx, ok := ptindex[f.p]; ok {
					tree[idx].me[pIdx] = uint8(f.e)
				}
			}
		}
		return
	}
	M := L + (R-L)/2
	if targetIdx <= M {
		update(2*idx, L, M, targetIdx, newValue)
	} else {
		update(2*idx+1, M+1, R, targetIdx, newValue)
	}
	tree[idx] = merge(tree[2*idx], tree[2*idx+1])
}

func query(idx, L, R, queryL, queryR int) Node {
	if queryL <= L && R <= queryR {
		return tree[idx]
	}
	if R < queryL || L > queryR {
		return Node{make([]uint8, K)}
	}
	M := L + (R-L)/2
	leftResult := query(2*idx, L, M, queryL, queryR)
	rightResult := query(2*idx+1, M+1, R, queryL, queryR)
	return merge(leftResult, rightResult)
}
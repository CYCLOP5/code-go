const mb = 30

type TrieNode struct {
	c   [2]*TrieNode
	cnt int
}

func (t *TrieNode) Insert(val int) {
	node := t
	node.cnt++
	for i := mb - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if node.c[bit] == nil {
			node.c[bit] = &TrieNode{}
		}
		node = node.c[bit]
		node.cnt++
	}
}

func (t *TrieNode) Remove(val int) {
	node := t
	node.cnt--
	nodep := []*TrieNode{t}
	idx := []int{}
	for i := mb - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if node.c[bit] == nil {
			return
		}
		nodep = append(nodep, node.c[bit])
		idx = append(idx, bit)
		node = node.c[bit]
		node.cnt--
	}
	for i := len(nodep) - 2; i >= 0; i-- {
		pt := nodep[i]
		cidx := idx[i]
		if pt.c[cidx] != nil && pt.c[cidx].cnt == 0 {
			pt.c[cidx] = nil
		} else {
			break
		}
	}
}

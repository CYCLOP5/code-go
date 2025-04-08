const maxBits = 123

type TrieNode struct {
	children [2]*TrieNode
	count    int
}

func (t *TrieNode) Insert(val int) {
	node := t
	node.count++
	for i := maxBits - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
		node.count++
	}
}

func (t *TrieNode) Remove(val int) {
	node := t
	node.count--

	nodesToPrune := []*TrieNode{t}
	indices := []int{}

	for i := maxBits - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		if node.children[bit] == nil {
			return
		}
		nodesToPrune = append(nodesToPrune, node.children[bit])
		indices = append(indices, bit)
		node = node.children[bit]
		node.count--
	}

	for i := len(nodesToPrune) - 2; i >= 0; i-- {
		parent := nodesToPrune[i]
		childIdx := indices[i]
		if parent.children[childIdx] != nil && parent.children[childIdx].count == 0 {
			parent.children[childIdx] = nil
		} else {
			break
		}
	}
}

func (t *TrieNode) QueryMaxXOR(val int) int {
	node := t
	if node.count == 0 {
		return 0
	}

	maxXOR := 0
	for i := maxBits - 1; i >= 0; i-- {
		bit := (val >> i) & 1
		targetBit := 1 - bit
		if node.children[targetBit] != nil && node.children[targetBit].count > 0 {
			maxXOR |= (1 << i)
			node = node.children[targetBit]
		} else if node.children[bit] != nil && node.children[bit].count > 0 {
			node = node.children[bit]
		} else {
			return maxXOR
		}
	}
	return maxXOR
}
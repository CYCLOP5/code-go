type TreapNode struct {
	key, priority int
	left, right   *TreapNode
}

func newNode(key int) *TreapNode {
	return &TreapNode{key: key, priority: rand.Int()}
}
func rotateRight(node *TreapNode) *TreapNode {
	l := node.left
	node.left = l.right
	l.right = node
	return l
}
func rotateLeft(node *TreapNode) *TreapNode {
	r := node.right
	node.right = r.left
	r.left = node
	return r
}
func insertTreap(node *TreapNode, key int) *TreapNode {
	if node == nil {
		return newNode(key)
	}
	if key < node.key {
		node.left = insertTreap(node.left, key)
		if node.left.priority > node.priority {
			node = rotateRight(node)
		}
	} else if key > node.key {
		node.right = insertTreap(node.right, key)
		if node.right.priority > node.priority {
			node = rotateLeft(node)
		}
	}
	return node
}
func deleteTreap(node *TreapNode, key int) *TreapNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = deleteTreap(node.left, key)
	} else if key > node.key {
		node.right = deleteTreap(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			if node.left.priority > node.right.priority {
				node = rotateRight(node)
				node.right = deleteTreap(node.right, key)
			} else {
				node = rotateLeft(node)
				node.left = deleteTreap(node.left, key)
			}
		}
	}
	return node
}
func findPredecessor(node *TreapNode, target int) (int, bool) {
	pred, ok := -1, false
	for node != nil {
		if node.key < target {
			pred, ok = node.key, true
			node = node.right
		} else {
			node = node.left
		}
	}
	return pred, ok
}
func findSuccessor(node *TreapNode, target int) (int, bool) {
	succ, ok := -1, false
	for node != nil {
		if node.key > target {
			succ, ok = node.key, true
			node = node.left
		} else {
			node = node.right
		}
	}
	return succ, ok
}

type Treap struct{ root *TreapNode }

func NewTreap() *Treap { return &Treap{} }
func (t *Treap) Contains(key int) bool {
	cur := t.root
	for cur != nil {
		if key == cur.key {
			return true
		}
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return false
}
func (t *Treap) Insert(key int) {
	if !t.Contains(key) {
		t.root = insertTreap(t.root, key)
	}
}
func (t *Treap) Delete(key int) {
	if t.Contains(key) {
		t.root = deleteTreap(t.root, key)
	}
}
func (t *Treap) FindPredecessor(x int) (int, bool) {
	return findPredecessor(t.root, x)
}
func (t *Treap) FindSuccessor(x int) (int, bool) {
	return findSuccessor(t.root, x)
}
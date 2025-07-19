type pqueue []*state
func (pq pqueue) Len() int { return len(pq) }
func (pq pqueue) Less(i, j int) bool {
    if pq[i].time != pq[j].time {
        return pq[i].time < pq[j].time
    }
    return pq[i].waits < pq[j].waits
}
func (pq pqueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *pqueue) Push(x interface{}) { *pq = append(*pq, x.(*state)) }
func (pq *pqueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[:n-1]
    return x
}
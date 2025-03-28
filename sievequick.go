var pms []int

var pfacs [ipmax + 1][]pair

func sieve() {
	ip := make([]bool, ipmax+1)
	for i := 2; i <= ipmax; i++ {
		ip[i] = true
	}
	ip[0], ip[1] = false, false
	for p := 2; p*p <= ipmax; p++ {
		if ip[p] {
			for i := p * p; i <= ipmax; i += p {
				ip[i] = false
			}
		}
	}
	for p := 2; p <= ipmax; p++ {
		if ip[p] {
			pms = append(pms, p)
		}
	}
}

func precomp() {
	for i := 1; i <= ipmax; i++ {
		n := i
		cf := []pair{}
		tn := n
		for _, p := range pms {
			if p*p > tn {
				break
			}
			if tn%p == 0 {
				cnt := 0
				for tn%p == 0 {
					tn /= p
					cnt++
				}
				cf = append(cf, pair{p, cnt})
			}
		}
		if tn > 1 {
			cf = append(cf, pair{tn, 1})
		}
		pfacs[i] = cf
	}
}

var K int                     
var ptindex map[int]int  

func initPrimes() {
	if len(pms) > 0 {
		return
	}
	sieve()
	precomp()
	K = len(pms)
	ptindex = make(map[int]int, K)
	for i, p := range pms {
		ptindex[p] = i
	}
}


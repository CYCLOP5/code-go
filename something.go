var scanner *FastReader
var out bytes.Buffer
const mod = 998244353
const tot = 500005

var fct [tot]int64
var ifct [tot]int64

func pow(b, pwr int64) int64 {
	rez := int64(1)
	b %= mod
	for pwr > 0 {
		if pwr%2 == 1 {
			rez = (rez * b) % mod
		}
		b = (b * b) % mod
		pwr /= 2
	}
	return rez
}

func inv(a int64) int64 {
	return pow(a, mod-2)
}

func prefac(maxN int) {
	if maxN <= 0 {
		if maxN == 0 {
			if len(fct) > 0 {
				fct[0] = 1
			}
			if len(ifct) > 0 {
				ifct[0] = 1
			}
		}
		return
	}

	fct[0] = 1
	ifct[0] = 1

	for i := 1; i < maxN; i++ {
		fct[i] = (fct[i-1] * int64(i)) % mod
	}

	if maxN > 1 {
		ifct[maxN-1] = inv(fct[maxN-1])
		for i := maxN - 2; i >= 1; i-- {
			ifct[i] = (ifct[i+1] * int64(i+1)) % mod
		}
	}
}
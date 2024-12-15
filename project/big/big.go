package big

import (
	gp "crycomproj/gopy"
	"crycomproj/util"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Success bool
	Time    time.Duration
	Message string
}

type Params struct {
	n, q, ex, m, N int
	qt             gp.T
	ell            int
	PO2            gp.Arr
}

type Big struct {
	Params
}

func (b *Big) Setup(q float64, n, m, ex int) {
	_, _, ell := util.GetQLogAndEll(q)

	b.q = int(q)
	b.qt = gp.T(q)
	b.n = n
	b.m = m
	b.ex = ex

	b.N = (b.n + 1) * ell
	b.PO2 = util.GeneratePowersOf2(ell, b.qt)

	b.ell = ell
}

func New(q float64, n, m, ex int) *Big {
	b := &Big{}
	b.Setup(q, n, m, ex)
	return b
}

func (b *Big) Run(iter int) chan Result {
	results := make(chan Result, iter)

	go func() {
		for i := 0; i < iter; i++ {
			result := Result{}
			start := time.Now()

			var mu gp.T = gp.T(rand.Intn(60))

			sk, t, v := SecretKeyGen(b.Params)
			pk := PublicKeyGen(b.Params, sk, t)

			C := Enc(b.Params, mu, pk)

			C = MultiplyConst(C, 10, b.n+1, b.ell, b.qt)

			mu_ := MPDec(b.Params, v, C)

			result.Message = fmt.Sprintf("%d vs %d", mu, mu_)
			if mu*10 == mu_ {
				result.Success = true
			}
			end := time.Now()
			result.Time = end.Sub(start)
			results <- result
		}
		close(results)
	}()

	return results
}

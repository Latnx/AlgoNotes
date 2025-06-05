package main

// https://leetcode-cn.com/problems/design-bitset/

type Bitset struct {
	x   []int
	cnt int
	n   int
}

func Constructor(size int) Bitset {
	return Bitset{
		x:   make([]int, (size+63)/64),
		cnt: 0,
		n:   size,
	}
}

func (this *Bitset) Fix(idx int) {
	if (this.x[idx/64]>>(idx%64))&1 == 0 {
		this.cnt++
	}
	this.x[idx/64] |= 1 << (idx % 64)
}

func (this *Bitset) Unfix(idx int) {
	if (this.x[idx/64]>>(idx%64))&1 == 1 {
		this.cnt--
	}
	this.x[idx/64] &= ^(1 << (idx % 64))
}

func (this *Bitset) Flip() {
	this.cnt = this.n - this.cnt
	for i, num := range this.x {
		this.x[i] = ^num
	}
}

func (this *Bitset) All() bool {
	return this.cnt == this.n
}

func (this *Bitset) One() bool {
	return this.cnt > 0
}

func (this *Bitset) Count() int {
	return this.cnt
}

func (this *Bitset) ToString() string {
	res := make([]byte, this.n)
	for i := range this.n {
		if (this.x[i/64]>>(i%64))&1 == 1 {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

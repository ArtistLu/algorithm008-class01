package main

func main() {

}

func findCircleNum(M [][]int) int {
	uf := newUf(M)
	for i := 0; i < len(M)-1; i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	return uf.count
}

type UF struct {
	parents []int
	count   int
	weitht  []int
}

func newUf(M [][]int) UF {
	count, weight, parents := len(M), make([]int, len(M)), make([]int, len(M))
	for i := 0; i < count; i++ {
		parents[i], weight[i] = i, 1
	}
	return UF{parents, count, weight}
}

func (uf *UF) find(i int) int {
	for i != uf.parents[i] {
		i, uf.parents[i] = uf.parents[i], uf.parents[uf.parents[i]]
	}
	return i
}

func (uf *UF) union(i, j int) {
	ip, jp := uf.find(i), uf.find(j)
	if ip == jp {
		return
	}

	if uf.weitht[ip] > uf.weitht[jp] {
		uf.parents[jp] = ip
		uf.weitht[ip] += uf.weitht[jp]
	} else {
		uf.parents[ip] = jp
		uf.weitht[jp] += uf.weitht[ip]
	}
	uf.count--
}

package main

import "fmt"

// ABC177 D
// union Find
// unionFindの構造体を作成
type UnionFind struct {
	parentSize []int //親ノードとサイズの配列
}

// コンストラクタみたいな役割
func NewUnionFind(n int) *UnionFind {
	parentSize := make([]int, n)
	for i := range parentSize {
		parentSize[i] = -1
	}
	return &UnionFind{parentSize}
}

// aとbを結合
func (uf *UnionFind) merge(a, b int) {
	x, y := uf.leader(a), uf.leader(b)
	if x == y {
		return
	}
	if abs(uf.parentSize[x]) < abs(uf.parentSize[y]) {
		x, y = y, x
	}
	uf.parentSize[x] += uf.parentSize[y]
	uf.parentSize[y] = x
}

// aとbが同じグループに属しているかを判定
func (uf *UnionFind) sama(a, b int) bool {
	return uf.leader(a) == uf.leader(b)
}

// aの親ノードを取得
func (uf *UnionFind) leader(a int) int {
	if uf.parentSize[a] < 0 {
		return a
	}
	uf.parentSize[a] = uf.leader(uf.parentSize[a])
	return uf.parentSize[a]
}

// aが属するグループのサイズを取得
func (uf *UnionFind) size(a int) int {
	return abs(uf.parentSize[uf.leader(a)])
}

// unionfindのグループを取得
func (uf *UnionFind) groups() [][]int {
	result := make([][]int, len(uf.parentSize))
	for i := range uf.parentSize {
		result[uf.leader(i)] = append(result[uf.leader(i)], i)
	}
	groups := make([][]int, 0)
	for _, r := range result {
		if len(r) > 0 {
			groups = append(groups, r)
		}
	}
	return groups
}

// 整数の絶対値
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	m := io.NextInt()

	uf := NewUnionFind(n)

	for i := 0; i < m; i++ {
		a := io.NextInt()
		b := io.NextInt()
		a--
		b--
		uf.merge(a, b)
	}

	friendsGroup := uf.groups()

	var maxFriendsSize int
	for _, fri := range friendsGroup {
		if len(fri) > maxFriendsSize {
			maxFriendsSize = len(fri)
		}
	}

	fmt.Println(maxFriendsSize)
}

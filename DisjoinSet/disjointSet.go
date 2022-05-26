package disjointSet

type DisjointSet interface {
	Find(a int) int
	Union(a, b int)
}

type disjointNode struct {
	id   int
	rank int
}

type disjointSet []disjointNode

func (ds *disjointSet) at(i int) *disjointNode {
	return &(*ds)[i]
}

func (ds *disjointSet) Find(a int) int {
	if a == ds.at(a).id {
		return a
	}
	ds.at(a).id = ds.Find((*ds)[a].id)
	return ds.at(a).id
}

func (ds *disjointSet) Union(a, b int) {
	a = ds.Find(a)
	b = ds.Find(b)
	if ds.at(a).rank < ds.at(b).rank {
		a, b = b, a
	}
	ds.at(b).id = a
	if ds.at(a).rank == ds.at(b).rank {
		ds.at(a).rank++
	}
}

func New(n int) *disjointSet {
	dsNodeArr := make([]disjointNode, n)
	ds := disjointSet(dsNodeArr)
	for i := 0; i < n; i++ {
		ds.at(i).id = i
		ds.at(i).rank = 0
	}
	return &ds
}

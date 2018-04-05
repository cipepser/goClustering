package ward

import (
	"math"

	"gonum.org/v1/gonum/floats"
)

// Group represents a cluster.
// Each Group is represented by `vec` and has N,
// the number of menber included in the group.
// Each Group have 2 children(left and right) excludes leaf node,
// leaf node's left and right is -1.
type Group struct {
	left, right, N int
	vec            []float64
	dist           float64
	visited        bool
}

// Tree represents a tree of groups.
type Tree []Group

// Ward execute Ward's method for clustering.
func Ward(X [][]float64) Tree {
	T := initialize(X)
	m := len(T)
	for {
		T = linkage(T)
		if len(T) == 2*m-1 {
			break
		}
	}
	return T
}

// linkage puts two groups into one.
func linkage(T Tree) Tree {
	d := math.Inf(1)
	idxi := len(T) + 1
	idxj := idxi

	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if !T[i].visited && !T[j].visited {
				dtmp := dist(T[i], T[j])
				if dtmp < d {
					d = dtmp
					idxi = i
					idxj = j
				}
			}
		}
	}

	vec := centerOfGravity(T[idxi], T[idxj])

	T = append(T, Group{
		left:    idxi,
		right:   idxj,
		N:       T[idxi].N + T[idxj].N,
		vec:     vec,
		dist:    d,
		visited: false,
	})
	T[idxi].visited = true
	T[idxj].visited = true

	return T
}

// centerOfGravity calcurates the gravity of two groups.
func centerOfGravity(g1, g2 Group) []float64 {
	vec := make([]float64, len(g1.vec))
	for i := 0; i < len(g1.vec); i++ {
		vec[i] = (float64(g1.N)*g1.vec[i] + float64(g2.N)*g2.vec[i]) / float64(g1.N+g2.N)
	}
	return vec
}

// initialize returns Tree consists of only leaf nodes.
func initialize(X [][]float64) Tree {
	T := make([]Group, len(X))
	for i, vec := range X {
		T[i] = Group{
			left:    -1,
			right:   -1,
			N:       1,
			vec:     vec,
			dist:    0,
			visited: false,
		}
	}
	return T
}

// dist calcurates distance between two groups.
func dist(g1, g2 Group) float64 {
	return float64(g1.N*g2.N) / float64(g1.N+g2.N) * math.Pow(floats.Distance(g1.vec, g2.vec, 2), 2)
}

// GetDist returns dist of the Group
func (g *Group) GetDist() float64 {
	return g.dist
}

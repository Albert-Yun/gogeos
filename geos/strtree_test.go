package geos

import (
	"testing"
)

func TestSTRTree(t *testing.T) {
	srtTree := SortedTileRecursiveTree(10)
	geom, _ := FromWKT(`POINT(1 1)`)
	srtTree.Insert(geom, "DADA")
	geom, _ = FromWKT(`POINT(2 2)`)
	srtTree.Insert(geom, "DIDI")
	geom, _ = FromWKT(`POINT(3 3)`)
	srtTree.Insert(geom, "DID3")

	g, _ := FromWKT(`POINT(2 2)`)
	srtTree.Query(g)
}

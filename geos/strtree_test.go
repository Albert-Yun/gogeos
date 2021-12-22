package geos

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSTRTree(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				a: `POINT(2 2)`,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				a: `POINT(12 12)`,
			},
			want: []int{0},
		},
		{
			args: args{
				a: `POINT(20 102)`,
			},
			want: make([]int, 0),
		},
	}

	geom, _ := FromWKT(`POLYGON(( 14 12, 14 14, 12 14, 12 12, 14 12))`)
	geom2, _ := FromWKT(`POLYGON(( 4 2, 4 4, 2 4, 2 2, 4 2))`)
	testData := []*Geometry{geom, geom2, geom2}
	strTree := SortedTileRecursiveTree(testData)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			result := make([]int, 0)
			g, _ := FromWKT(tt.args.a)
			strTree.Query(g, func(i int) {
				result = append(result, i)
			})

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("should equal! got %v want %v", result, tt.want)
			}
		})
	}
}

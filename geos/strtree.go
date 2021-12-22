package geos

/*
#include "geos.h"
extern void go_callback(void* item, void* user_data);
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

type STRTree struct {
	t *C.GEOSSTRtree
}

func (t *STRTree) destroy() {
	cGEOSSTRtree_destroy(t.t)
	t.t = nil
}

func SortedTileRecursiveTree(geom []*Geometry) *STRTree {
	csize := C.size_t(len(geom))
	t := cGEOSSTRtree_create(csize)
	if t == nil {
		return nil
	}

	tree := &STRTree{
		t: t,
	}

	for i, g := range geom {
		tree.insert(g, i)
	}

	runtime.SetFinalizer(tree, (*STRTree).destroy)
	return tree
}

func (t *STRTree) insert(geom *Geometry, item int) {
	cIdx := C.int(item)
	cGEOSSTRtree_insert(t.t, geom.g, (*C.void)(unsafe.Pointer(&cIdx)))
}

func (t *STRTree) Query(geom *Geometry, cb STRTreeCallback)  {
	cbid := register(cb)
	ccbid := C.int(cbid)
	cGEOSSTRtree_query(t.t, geom.g, (C.GEOSQueryCallback)(C.go_callback), (*C.void)(unsafe.Pointer(&ccbid)))
	defer unregister(cbid)
}

// cgo callback
// https://github.com/golang/go/wiki/cgo#function-variables
//export go_callback
func go_callback(cIdx unsafe.Pointer, data unsafe.Pointer) {
	goIdx := *(*int)(cIdx)
	cbid := *(*C.int)(data)
	cb := lookup(int(cbid))
	cb(goIdx)
}

type STRTreeCallback func(id int)

var mu sync.Mutex
var fns = make(map[int]STRTreeCallback)
var cbIndex int

func register(fn STRTreeCallback) int {
	mu.Lock()
	defer mu.Unlock()
	cbIndex++

	for fns[cbIndex] != nil {
		cbIndex++
	}

	fns[cbIndex] = fn
	return cbIndex
}

func lookup(i int) STRTreeCallback {
	mu.Lock()
	defer mu.Unlock()
	return fns[i]
}

func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, i)
}

//func cGEOSSTRtree_nearest_generic(tree *C.GEOSSTRtree, item *C.void, itemEnvelope *C.GEOSGeometry, distancefn C.GEOSDistanceCallback, userdata *C.void) {
//	handlemu.Lock()
//	defer handlemu.Unlock()
//	C.GEOSSTRtree_nearest_generic_r(handle, tree, unsafe.Pointer(item), itemEnvelope, distancefn, unsafe.Pointer(userdata))
//}
//
//func cGEOSSTRtree_iterate(tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
//	handlemu.Lock()
//	defer handlemu.Unlock()
//	C.GEOSSTRtree_iterate_r(handle, tree, callback, unsafe.Pointer(userdata))
//}
//
//func cGEOSSTRtree_remove(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
//	handlemu.Lock()
//	defer handlemu.Unlock()
//	return C.GEOSSTRtree_remove_r(handle, tree, g, unsafe.Pointer(item))
//}

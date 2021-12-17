package geos

/*
#include "geos.h"
extern void go_callback(void* item, void* user_data);
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type STRTree struct {
	t *C.GEOSSTRtree
}

func (t *STRTree) destroy() {
	cGEOSSTRtree_destroy(t.t)
	t.t = nil
}

func SortedTileRecursiveTree(size int) *STRTree {
	csize := C.size_t(size)
	t := cGEOSSTRtree_create(csize)
	if t == nil {
		return nil
	}

	tree := &STRTree{
		t: t,
	}

	runtime.SetFinalizer(tree, (*STRTree).destroy)
	return tree
}

func (t *STRTree) Insert(geom *Geometry, item string) {
	cstr := C.CString(item)
	defer C.free(unsafe.Pointer(cstr))

	cGEOSSTRtree_insert(t.t, geom.g, (*C.void)(unsafe.Pointer(cstr)))
}

//export go_callback
func go_callback(item *C.void, userData *C.void) {
	fmt.Printf("%v %v\n", item, userData)
}

func (t *STRTree) Query(geom *Geometry) {
	cstr := C.CString("DIDI")
	defer C.free(unsafe.Pointer(cstr))

	cb := (C.GEOSQueryCallback)(C.go_callback)
	cGEOSSTRtree_query(t.t, geom.g, cb, (*C.void)(unsafe.Pointer(cstr)))
}

//
//func cGEOSSTRtree_query(tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
//	handlemu.Lock()
//	defer handlemu.Unlock()
//	C.GEOSSTRtree_query_r(handle, tree, g, callback, unsafe.Pointer(userdata))
//}
//
//func cGEOSSTRtree_nearest(tree *C.GEOSSTRtree, geom *C.GEOSGeometry) *C.GEOSGeometry {
//	handlemu.Lock()
//	defer handlemu.Unlock()
//	return C.GEOSSTRtree_nearest_r(handle, tree, geom)
//}
//
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

package rtda

import "miniJVM/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

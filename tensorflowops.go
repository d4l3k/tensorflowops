package tensorflowops

// #include <stdlib.h>
// #include "tensorflow/c/c_api.h"
// #cgo LDFLAGS: -ltensorflow
import "C"

import (
	"unsafe"

	_ "github.com/tensorflow/tensorflow/tensorflow/go"

	"github.com/d4l3k/tensorflowops/proto/github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
	"github.com/gogo/protobuf/proto"
)

func RegisteredOps() ([]*framework.OpDef, error) {
	buf := C.TF_GetAllOpList()
	defer C.TF_DeleteBuffer(buf)
	var (
		list framework.OpList
		size = int(buf.length)
		// A []byte backed by C memory.
		// See: https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
		data = (*[1 << 30]byte)(unsafe.Pointer(buf.data))[:size:size]
		err  = proto.Unmarshal(data, &list)
	)
	if err != nil {
		return nil, err
	}
	return list.GetOp(), nil
}

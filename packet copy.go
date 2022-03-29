// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

func toCPacket(pkt *Packet) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *Packet {
	return (*Packet)(unsafe.Pointer(pkt))
}

// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

func (p *AVPacket) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}
func (p *AVPacket) Duration() int {
	return int(p.duration)
}

func (p *AVPacket) SetDuration(d int64) {
	p.duration = C.long(d)
}

func (p *AVPacket) Flags() int {
	return int(p.flags)
}
func (p *AVPacket) SetFlags(flags int) {
	p.flags = C.int(flags)
}
func (p *AVPacket) SideDataElems() int {
	return int(p.side_data_elems)
}
func (p *AVPacket) Size() int {
	return int(p.size)
}
func (p *AVPacket) StreamIndex() int {
	return int(p.stream_index)
}
func (p *AVPacket) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}
func (p *AVPacket) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}
func (p *AVPacket) Dts() int64 {
	return int64(p.dts)
}
func (p *AVPacket) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}
func (p *AVPacket) Pos() int64 {
	return int64(p.pos)
}
func (p *AVPacket) Pts() int64 {
	return int64(p.pts)
}
func (p *AVPacket) SetPts(pts int64) {
	p.pts = C.int64_t(pts)
}
func (p *AVPacket) Data() *uint8 {
	return (*uint8)(p.data)
}

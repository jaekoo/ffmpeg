// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *Parser) AvParserNext() *Parser {
	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *Parser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

func (ctxt *Context) SetTimebase(num1 int, den1 int) {
	ctxt.time_base.num = C.int(num1)
	ctxt.time_base.den = C.int(den1)
}

func (ctxt *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctxt.width = C.int(width)
	ctxt.height = C.int(height)
	// ctxt.bit_rate = 1000000
	ctxt.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	if hasBframes {
		ctxt.has_b_frames = 1
	} else {
		ctxt.has_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

func (ctxt *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctxt.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (ctxt *Context) AvcodecSendPacket(packet *AVPacket) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

func (ctxt *Context) AvcodecReceiveFrame(frame *AVFrame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

func (ctxt *Context) AvcodecReceivePacket(packet *AVPacket) int {
	return (int)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

func (ctxt *Context) AvcodecSendFrame(frame *AVFrame) int {
	return (int)(C.avcodec_send_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

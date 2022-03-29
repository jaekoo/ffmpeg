// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//Rational av_stream_get_r_frame_rate (const AVStream *s)
func (s *AVStream) AvStreamGetRFrameRate() Rational {
	return newRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (AVStream *s, Rational r)
func (s *AVStream) AvStreamSetRFrameRate(r Rational) {
	rat := C.struct_AVRational{
		num: C.int(r.Num()),
		den: C.int(r.Den()),
	}

	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), rat)
}

//struct CodecParserContext * av_stream_get_parser (const AVStream *s)
func (s *AVStream) AvStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// //char * av_stream_get_recommended_encoder_configuration (const AVStream *s)
// func (s *AVStream) AvStreamGetRecommendedEncoderConfiguration() string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (AVStream *s, char *configuration)
// func (s *AVStream) AvStreamSetRecommendedEncoderConfiguration( c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const AVStream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *AVStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

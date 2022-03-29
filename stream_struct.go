// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

type (
	CodecParserContext C.struct_AVCodecParserContext
	AvIndexEntry       C.struct_AVIndexEntry
	AvStreamParseType  C.enum_AVStreamParseType
	AvPacketList       C.struct_AVPacketList
	AVStream           C.struct_AVStream
)

func (avs *AVStream) CodecParameters() *AvCodecParameters {
	return (*AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

func (avs *AVStream) Codec() *Context {
	return (*Context)(unsafe.Pointer(avs.codec))
}

func (avs *AVStream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(avs.metadata))
}

func (avs *AVStream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

func (avs *AVStream) AttachedPic() Packet {
	return *fromCPacket(&avs.attached_pic)
}

func (avs *AVStream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

func (avs *AVStream) ProbeData() AvProbeData {
	return AvProbeData(avs.probe_data)
}

func (avs *AVStream) AvgFrameRate() Rational {
	return newRational(avs.avg_frame_rate)
}

// func (avs *AVStream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *AVStream) RFrameRate() Rational {
	return newRational(avs.r_frame_rate)
}

func (avs *AVStream) SampleAspectRatio() Rational {
	return newRational(avs.sample_aspect_ratio)
}

func (avs *AVStream) TimeBase() Rational {
	return newRational(avs.time_base)
}

// func (avs *AVStream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *AVStream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

func (avs *AVStream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

func (avs *AVStream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *AVStream) Disposition() int {
	return int(avs.disposition)
}

func (avs *AVStream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *AVStream) Id() int {
	return int(avs.id)
}

func (avs *AVStream) SetID(id int32) {
	//avs.id = _C_int(id)
}

func (avs *AVStream) Index() int {
	return int(avs.index)
}

func (avs *AVStream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

func (avs *AVStream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

func (avs *AVStream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

func (avs *AVStream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

func (avs *AVStream) NbSideData() int {
	return int(avs.nb_side_data)
}

func (avs *AVStream) ProbePackets() int {
	return int(avs.probe_packets)
}

func (avs *AVStream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

func (avs *AVStream) RequestProbe() int {
	return int(avs.request_probe)
}

func (avs *AVStream) SkipSamples() int {
	return int(avs.skip_samples)
}

func (avs *AVStream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

func (avs *AVStream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

func (avs *AVStream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

func (avs *AVStream) CurDts() int64 {
	return int64(avs.cur_dts)
}

func (avs *AVStream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *AVStream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

func (avs *AVStream) FirstDts() int64 {
	return int64(avs.first_dts)
}

func (avs *AVStream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

func (avs *AVStream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *AVStream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (avs *AVStream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

func (avs *AVStream) LastIpPts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *AVStream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

func (avs *AVStream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

func (avs *AVStream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

func (avs *AVStream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

func (avs *AVStream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *AVStream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *AVStream) StartTime() int64 {
	return int64(avs.start_time)
}

func (avs *AVStream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.parser))
}

func (avs *AVStream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *AVStream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

func (avs *AVStream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

func (avs *AVStream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

func (avs *AVStream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

func (avs *AVStream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

func (avs *AVStream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}

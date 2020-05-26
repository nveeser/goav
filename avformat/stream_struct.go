// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.ptr.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/nveeser/goav/avcodec"
	"github.com/nveeser/goav/avutil"
)

func (s *Stream) CodecParameters() *avcodec.AvCodecParameters {
	return (*avcodec.AvCodecParameters)(unsafe.Pointer(s.ptr.codecpar))
}

func (s *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(s.ptr.codec))
}

func (s *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(s.ptr.metadata))
}

func (s *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(s.ptr.index_entries))
}

func (s *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&s.ptr.attached_pic)
}

func (s *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(s.ptr.side_data))
}

func (s *Stream) ProbeData() AvProbeData {
	return AvProbeData(s.ptr.probe_data)
}

func (s *Stream) AvgFrameRate() avcodec.Rational {
	return newRational(s.ptr.avg_frame_rate)
}

// func (avs *Stream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.ptr.display_aspect_ratio))
// }

func (s *Stream) RFrameRate() avcodec.Rational {
	return newRational(s.ptr.r_frame_rate)
}

func (s *Stream) SampleAspectRatio() avcodec.Rational {
	return newRational(s.ptr.sample_aspect_ratio)
}

func (s *Stream) TimeBase() avcodec.Rational {
	return newRational(s.ptr.time_base)
}

// func (avs *Stream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.ptr.recommended_encoder_configuration)
// }

func (s *Stream) Discard() AvDiscard {
	return AvDiscard(s.ptr.discard)
}

func (s *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(s.ptr.need_parsing)
}

func (s *Stream) CodecInfoNbFrames() int {
	return int(s.ptr.codec_info_nb_frames)
}

func (s *Stream) Disposition() int {
	return int(s.ptr.disposition)
}

func (s *Stream) EventFlags() int {
	return int(s.ptr.event_flags)
}

func (s *Stream) Id() int {
	return int(s.ptr.id)
}

func (s *Stream) Index() int {
	return int(s.ptr.index)
}

func (s *Stream) InjectGlobalSideData() int {
	return int(s.ptr.inject_global_side_data)
}

func (s *Stream) LastIpDuration() int {
	return int(s.ptr.last_IP_duration)
}

func (s *Stream) NbDecodedFrames() int {
	return int(s.ptr.nb_decoded_frames)
}

func (s *Stream) NbIndexEntries() int {
	return int(s.ptr.nb_index_entries)
}

func (s *Stream) NbSideData() int {
	return int(s.ptr.nb_side_data)
}

func (s *Stream) ProbePackets() int {
	return int(s.ptr.probe_packets)
}

func (s *Stream) PtsWrapBehavior() int {
	return int(s.ptr.pts_wrap_behavior)
}

func (s *Stream) RequestProbe() int {
	return int(s.ptr.request_probe)
}

func (s *Stream) SkipSamples() int {
	return int(s.ptr.skip_samples)
}

func (s *Stream) SkipToKeyframe() int {
	return int(s.ptr.skip_to_keyframe)
}

func (s *Stream) StreamIdentifier() int {
	return int(s.ptr.stream_identifier)
}

func (s *Stream) UpdateInitialDurationsDone() int {
	return int(s.ptr.update_initial_durations_done)
}

func (s *Stream) CurDts() int64 {
	return int64(s.ptr.cur_dts)
}

func (s *Stream) Duration() int64 {
	return int64(s.ptr.duration)
}

// func (avs *Stream) FirstDiscardSample() int64 {
// 	return int64(avs.ptr.first_discard_sample)
// }

func (s *Stream) FirstDts() int64 {
	return int64(s.ptr.first_dts)
}

func (s *Stream) InterleaverChunkDuration() int64 {
	return int64(s.ptr.interleaver_chunk_duration)
}

func (s *Stream) InterleaverChunkSize() int64 {
	return int64(s.ptr.interleaver_chunk_size)
}

// func (avs *Stream) LastDiscardSample() int64 {
// 	return int64(avs.ptr.last_discard_sample)
// }

func (s *Stream) LastDtsForOrderCheck() int64 {
	return int64(s.ptr.last_dts_for_order_check)
}

func (s *Stream) LastIpPts() int64 {
	return int64(s.ptr.last_IP_pts)
}

func (s *Stream) MuxTsOffset() int64 {
	return int64(s.ptr.mux_ts_offset)
}

func (s *Stream) NbFrames() int64 {
	return int64(s.ptr.nb_frames)
}

func (s *Stream) PtsBuffer() int64 {
	return int64(s.ptr.pts_buffer[0])
}

func (s *Stream) PtsReorderError() int64 {
	return int64(s.ptr.pts_reorder_error[0])
}

func (s *Stream) PtsWrapReference() int64 {
	return int64(s.ptr.pts_wrap_reference)
}

// func (avs *Stream) StartSkipSamples() int64 {
// 	return int64(avs.ptr.start_skip_samples)
// }

func (s *Stream) StartTime() int64 {
	return int64(s.ptr.start_time)
}

func (s *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(s.ptr.parser))
}

func (s *Stream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(s.ptr.last_in_packet_buffer))
}

// func (avs *Stream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.ptr.priv_pts))
// }

func (s *Stream) DtsMisordered() uint8 {
	return uint8(s.ptr.dts_misordered)
}

func (s *Stream) DtsOrdered() uint8 {
	return uint8(s.ptr.dts_ordered)
}

func (s *Stream) PtsReorderErrorCount() uint8 {
	return uint8(s.ptr.pts_reorder_error_count[0])
}

func (s *Stream) IndexEntriesAllocatedSize() uint {
	return uint(s.ptr.index_entries_allocated_size)
}

func (s *Stream) Free() {
	C.av_freep(unsafe.Pointer(s))
}

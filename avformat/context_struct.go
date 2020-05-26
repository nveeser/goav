// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/nveeser/goav/avutil"
)

func (c *Context) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(c.ptr.chapters))
}

func (c *Context) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.ptr.audio_codec))
}

func (c *Context) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.ptr.subtitle_codec))
}

func (c *Context) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(c.ptr.video_codec))
}

func (c *Context) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(c.ptr.metadata))
}

func (c *Context) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(c.ptr.internal))
}

func (c *Context) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(c.ptr.pb))
}

func (c *Context) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(c.ptr.interrupt_callback)
}

func (c *Context) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c.ptr.programs)),
		Len:  int(c.NbPrograms()),
		Cap:  int(c.NbPrograms()),
	}
	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

func (c *Context) Streams() []*Stream {
	var cstream []*C.struct_AVStream
	slice := (*reflect.SliceHeader)((unsafe.Pointer(&cstream)))
	slice.Cap = int(c.ptr.nb_streams)
	slice.Len = int(c.ptr.nb_streams)
	slice.Data = uintptr(unsafe.Pointer(c.ptr.streams))

	streams := make([]*Stream, len(cstream))
	for i, cs := range cstream {
		streams[i] = &Stream{
			context: c,
			ptr:     cs,
		}
	}
	return streams
}

func (c *Context) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&c.ptr.filename[0])))
}

// func (c *Context) CodecWhitelist() string {
// 	return C.GoString(c.ptr.codec_whitelist)
// }

// func (c *Context) FormatWhitelist() string {
// 	return C.GoString(c.ptr.format_whitelist)
// }

func (c *Context) AudioCodecId() CodecId {
	return CodecId(c.ptr.audio_codec_id)
}

func (c *Context) SubtitleCodecId() CodecId {
	return CodecId(c.ptr.subtitle_codec_id)
}

func (c *Context) VideoCodecId() CodecId {
	return CodecId(c.ptr.video_codec_id)
}

func (c *Context) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(c.ptr.duration_estimation_method)
}

func (c *Context) AudioPreload() int {
	return int(c.ptr.audio_preload)
}

func (c *Context) AvioFlags() int {
	return int(c.ptr.avio_flags)
}

func (c *Context) AvoidNegativeTs() int {
	return int(c.ptr.avoid_negative_ts)
}

func (c *Context) BitRate() int {
	return int(c.ptr.bit_rate)
}

func (c *Context) CtxFlags() int {
	return int(c.ptr.ctx_flags)
}

func (c *Context) Debug() int {
	return int(c.ptr.debug)
}

func (c *Context) ErrorRecognition() int {
	return int(c.ptr.error_recognition)
}

func (c *Context) EventFlags() int {
	return int(c.ptr.event_flags)
}

func (c *Context) Flags() int {
	return int(c.ptr.flags)
}

func (c *Context) FlushPackets() int {
	return int(c.ptr.flush_packets)
}

func (c *Context) FormatProbesize() int {
	return int(c.ptr.format_probesize)
}

func (c *Context) FpsProbeSize() int {
	return int(c.ptr.fps_probe_size)
}

func (c *Context) IoRepositioned() int {
	return int(c.ptr.io_repositioned)
}

func (c *Context) Keylen() int {
	return int(c.ptr.keylen)
}

func (c *Context) MaxChunkDuration() int {
	return int(c.ptr.max_chunk_duration)
}

func (c *Context) MaxChunkSize() int {
	return int(c.ptr.max_chunk_size)
}

func (c *Context) MaxDelay() int {
	return int(c.ptr.max_delay)
}

func (c *Context) MaxTsProbe() int {
	return int(c.ptr.max_ts_probe)
}

func (c *Context) MetadataHeaderPadding() int {
	return int(c.ptr.metadata_header_padding)
}

func (c *Context) ProbeScore() int {
	return int(c.ptr.probe_score)
}

func (c *Context) Seek2any() int {
	return int(c.ptr.seek2any)
}

func (c *Context) StrictStdCompliance() int {
	return int(c.ptr.strict_std_compliance)
}

func (c *Context) TsId() int {
	return int(c.ptr.ts_id)
}

func (c *Context) UseWallclockAsTimestamps() int {
	return int(c.ptr.use_wallclock_as_timestamps)
}

func (c *Context) Duration() int64 {
	return int64(c.ptr.duration)
}

func (c *Context) MaxAnalyzeDuration2() int64 {
	return int64(c.ptr.max_analyze_duration)
}

func (c *Context) MaxInterleaveDelta() int64 {
	return int64(c.ptr.max_interleave_delta)
}

func (c *Context) OutputTsOffset() int64 {
	return int64(c.ptr.output_ts_offset)
}

func (c *Context) Probesize2() int64 {
	return int64(c.ptr.probesize)
}

func (c *Context) SkipInitialBytes() int64 {
	return int64(c.ptr.skip_initial_bytes)
}

func (c *Context) StartTime() int64 {
	return int64(c.ptr.start_time)
}

func (c *Context) StartTimeRealtime() int64 {
	return int64(c.ptr.start_time_realtime)
}

func (c *Context) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(c.ptr.iformat))
}

func (c *Context) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(c.ptr.oformat))
}

// func (c *Context) DumpSeparator() uint8 {
// 	return uint8(c.ptr.dump_separator)
// }

func (c *Context) CorrectTsOverflow() int {
	return int(c.ptr.correct_ts_overflow)
}

func (c *Context) MaxIndexSize() uint {
	return uint(c.ptr.max_index_size)
}

func (c *Context) MaxPictureBuffer() uint {
	return uint(c.ptr.max_picture_buffer)
}

func (c *Context) NbChapters() uint {
	return uint(c.ptr.nb_chapters)
}

func (c *Context) NbPrograms() uint {
	return uint(c.ptr.nb_programs)
}

func (c *Context) NbStreams() uint {
	return uint(c.ptr.nb_streams)
}

func (c *Context) PacketSize() uint {
	return uint(c.ptr.packet_size)
}

func (c *Context) Probesize() uint {
	return uint(c.ptr.probesize)
}

func (c *Context) SetPb(pb *AvIOContext) {
	c.ptr.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

func (c *Context) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&c.ptr.pb))
}

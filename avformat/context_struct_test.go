package avformat

import (
	"runtime"
	"testing"
)

func TestFinalizer(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		ctx := NewContext()
		_ = ctx
	}
	runtime.GC()

	got := allocatedContextCount()
	if got == 1000 {
		t.Errorf("allocatedContextCount got %d == 1000 wanted < 1000 ", got)
	}
}

func TestContextStruct(t *testing.T) {
	ctx := NewContext()
	if err := ctx.OpenInput("small.mp4", nil, nil); err != nil {
		t.Errorf("OpenInput got %q wanted nil", err)
	}
	ctx.Filename()

	ctx.AudioCodecId()
	ctx.SubtitleCodecId()
	ctx.VideoCodecId()
	ctx.AudioPreload()
	ctx.AvioFlags()
	ctx.AvoidNegativeTs()
	ctx.BitRate()
	ctx.CtxFlags()
	ctx.Debug()
	ctx.ErrorRecognition()
	ctx.EventFlags()
	ctx.Flags()
	ctx.FlushPackets()
	ctx.FormatProbesize()
	ctx.FpsProbeSize()
	ctx.IoRepositioned()
	ctx.Keylen()
	ctx.MaxChunkDuration()
	ctx.MaxChunkSize()
	ctx.MaxDelay()
	ctx.MaxTsProbe()
	ctx.MetadataHeaderPadding()
	ctx.ProbeScore()
	ctx.Seek2any()
	ctx.StrictStdCompliance()
	ctx.TsId()
	ctx.UseWallclockAsTimestamps()
	ctx.Duration()
	ctx.MaxAnalyzeDuration2()
	ctx.MaxInterleaveDelta()
	ctx.OutputTsOffset()
	ctx.Probesize2()
	ctx.SkipInitialBytes()
	ctx.StartTime()
	ctx.StartTimeRealtime()
	ctx.CorrectTsOverflow()
	ctx.MaxIndexSize()
	ctx.MaxPictureBuffer()
	ctx.NbChapters()
	ctx.NbPrograms()
	ctx.NbStreams()
	ctx.PacketSize()
	ctx.Probesize()

}

package avformat

import (
	"syscall"
	"testing"
)

func TestFinalizer(t*testing.T) {
	for sample := 1; sample <= sampleSize; sample++ {
		for i := 0; i < loopSize; i++ {
			a := NewAllocator(10 * 1024)
			_ = a
		}

		err := syscall.Getrusage(syscall.RUSAGE_SELF, &rUsage)
		if err != nil {
			panic(err)
		}
	}
}

func TestContextStruct(t *testing.T) {
	ctx := NewContext()
	if err := ctx.OpenInput("small.mp4", nil, nil); err != nil{
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

package main

import (
	"log"

	"github.com/nveeser/goav/avcodec"
	"github.com/nveeser/goav/avdevice"
	"github.com/nveeser/goav/avfilter"
	"github.com/nveeser/goav/avformat"
	"github.com/nveeser/goav/avutil"
	"github.com/nveeser/goav/swresample"
	"github.com/nveeser/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}

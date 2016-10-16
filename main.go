package main

/*

#cgo pkg-config: libavformat libavdevice

#include <stdlib.h>
#include "libavformat/avformat.h"
#include <libavdevice/avdevice.h>
#include "c/segmenter.h"

*/
import "C"
import "fmt"

type config struct {
	BaseURL         string
	FileBase        string
	MediaFileName   string
	IndexFile       string
	SourceFile      string
	Media           int
	Stat            int
	PlaylistEntries int
	Duration        int
}

func main() {
	var avContext *C.struct_AVFormatContext
	//	var segmenterContext *C.struct_SegmenterContext

	C.av_register_all()
	C.avformat_free_context(avContext)
	fmt.Println("it works.")
}

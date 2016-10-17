package main

/*
#include <stdio.h>
#include "libavformat/avformat.h"
#include <libavdevice/avdevice.h>
#include "c/segmenter.h"
#include "c/util.h"

#cgo LDFLAGS: -L${SRCDIR}/build -lsegmenter -lavcodec -lavformat -lavutil
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/3d0c/gmf"
)

type config struct {
	BaseURL         string
	FileBase        string
	MediaFileName   string
	IndexFile       string
	SourceFile      string
	Stat            int
	PlaylistEntries int
	Duration        float64
}

func segment(cfg config) error {
	var sourceContext *C.struct_AVFormatContext
	var outputContext *C.SegmenterContext

	if cfg.MediaBaseName == "" {
		cfg.MediaBaseName = "fileSequence"
	}

	if cfg.Duration == 0 {
		cfg.Duration = 10
	}

	if cfg.IndexFile == "" {
		cfg.IndexFile = "prog_index.m3u8"
	}

	//TODO improve the way we create the folder
	os.Mkdir(cfg.FileBase, 0777)

	C.av_register_all()

	sourceFile := C.CString(cfg.SourceFile)
	defer C.free(unsafe.Pointer(sourceFile))

	if averr := C.avformat_open_input(&sourceContext, sourceFile, nil, nil); averr < 0 {
		return fmt.Errorf("Error opening input: %s", gmf.AvError(int(averr)))
	}
	defer C.avformat_free_context(sourceContext)

	if averr := C.avformat_find_stream_info(sourceContext, nil); averr < 0 {
		return fmt.Errorf("Error finding stream info: %s", gmf.AvError(int(averr)))
	}

	if ret := C.segmenter_alloc_context(&outputContext); ret < 0 {
		return fmt.Errorf("Cannot allocate context: %s", C.sg_strerror(ret))
	}
	defer C.segmenter_free_context(outputContext)

	return nil
}

func main() {
	myCfg := config{
		SourceFile: "fixtures/test.mp4",
	}

	if res := segment(myCfg); res != nil {
		fmt.Println("error!", res.Error())
	} else {
		fmt.Println("success")
	}
}

package main

/*

#cgo pkg-config: libavformat libavdevice

#include <stdlib.h>
#include "libavformat/avformat.h"
#include <libavdevice/avdevice.h>
#include "c/segmenter.h"

*/
import "C"

import (
	"fmt"

	"github.com/3d0c/gmf"
)

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

func segment(cfg config) error {
	var sourceContext *C.struct_AVFormatContext
	//	var segmenterContext *C.struct_SegmenterContext

	C.av_register_all()

	if averr := C.avformat_open_input(&sourceContext, C.CString(cfg.SourceFile), nil, nil); averr < 0 {
		return fmt.Errorf("Error opening input: %s", gmf.AvError(int(averr)))
	}

	C.avformat_free_context(sourceContext)
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

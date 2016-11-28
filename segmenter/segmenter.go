package segmenter

/*
#include <stdio.h>
#include "libavformat/avformat.h"
#include <libavdevice/avdevice.h>
#include "c/segmenter.h"
#include "c/util.h"

#cgo LDFLAGS: -L${SRCDIR}/../build -lsegmenter -lavcodec -lavformat -lavutil
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/3d0c/gmf"
)

//HLSConfig stores all the configurations needed for HLS outputs
type HLSConfig struct {
	BaseURL       string
	FileBase      string
	MediaBaseName string
	IndexFile     string
	SourceFile    string
	Stat          int
	Duration      float64
}

// Segment function is responsible for the HLS generation
func Segment(cfg HLSConfig) error {
	var sourceContext *C.struct_AVFormatContext
	var outputContext *C.SegmenterContext

	cfg = setDefaultValues(cfg)

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
		return fmt.Errorf("Cannot allocate context: %s", C.GoString(C.sg_strerror(C.int(ret))))
	}
	defer C.segmenter_free_context(outputContext)

	fileBase := C.CString(cfg.FileBase)
	defer C.free(unsafe.Pointer(fileBase))

	mediaBaseName := C.CString(cfg.MediaBaseName)
	defer C.free(unsafe.Pointer(mediaBaseName))

	if ret := C.segmenter_init(outputContext, sourceContext, fileBase, mediaBaseName, C.double(cfg.Duration), C.int(3)); ret != 0 {
		return fmt.Errorf("Cannot initialize segmenter: %s", C.GoString(C.sg_strerror(C.int(ret))))
	}

	if ret := C.segmenter_open(outputContext); ret < 0 {
		return fmt.Errorf("open output: %s", C.GoString(C.sg_strerror(C.int(ret))))
	}

	var pkt C.AVPacket

	for {
		if ret := C.av_read_frame(sourceContext, &pkt); ret < 0 {
			break
		}

		if ret := C.segmenter_write_pkt(outputContext, sourceContext, &pkt); ret != 0 {
			return fmt.Errorf("writing packet: %s", C.GoString(C.sg_strerror(C.int(ret))))
		}
	}

	C.segmenter_close(outputContext)

	baseURL := C.CString(cfg.BaseURL)
	defer C.free(unsafe.Pointer(baseURL))

	indexFile := C.CString(cfg.IndexFile)
	defer C.free(unsafe.Pointer(indexFile))

	if ret := C.segmenter_write_playlist(outputContext, 0, baseURL, indexFile); ret != 0 {
		return fmt.Errorf("writing playlist: %s", C.GoString(C.sg_strerror(C.int(ret))))
	}

	return nil
}

func setDefaultValues(cfg HLSConfig) HLSConfig {
	if cfg.MediaBaseName == "" {
		cfg.MediaBaseName = "segment"
	}

	if cfg.Duration == 0 {
		cfg.Duration = 10
	}

	if cfg.IndexFile == "" {
		cfg.IndexFile = "video.m3u8"
	}

	return cfg
}

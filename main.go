package main

import "fmt"
import "github.com/snickers/hls/segmenter"

func main() {
	myCfg := segmenter.HLSConfig{
		SourceFile: "fixtures/test.mp4",
		FileBase:   "output",
	}

	if res := segmenter.Segment(myCfg); res != nil {
		fmt.Println("error!", res.Error())
	} else {
		fmt.Println("success")
	}
}

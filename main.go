package main

import "fmt"

func main() {
	myCfg := HLSConfig{
		SourceFile: "fixtures/test.mp4",
		FileBase:   "output",
	}

	if res := Segment(myCfg); res != nil {
		fmt.Println("error!", res.Error())
	} else {
		fmt.Println("success")
	}
}

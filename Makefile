all:
	rm -rf build libsegmenter.so
	mkdir build
	$(CC) -Wno-deprecated -g -fPIC -c -o build/segmenter.o c/segmenter.c
	$(CC) -Wno-deprecated -g -fPIC -c -o build/util.o c/util.c
	$(CC) -g -lavcodec -lavformat -lavutil -fPIC -shared -o build/libsegmenter.so build/segmenter.o build/util.o
	LD_LIBRARY_PATH=. go build main.go

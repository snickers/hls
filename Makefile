all: clean dep
	go build

clean:
	@rm -rf build hls

dep:
	@mkdir build
	@$(CC) -Wno-deprecated -g -fPIC -c -o build/segmenter.o segmenter/c/segmenter.c
	@$(CC) -Wno-deprecated -g -fPIC -c -o build/util.o segmenter/c/util.c
	@$(CC) -g -lavcodec -lavformat -lavutil -fPIC -shared -o build/libsegmenter.so build/segmenter.o build/util.o



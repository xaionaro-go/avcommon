deps: 3rdparty/ffmpeg/mediacodec_wrapper.h 3rdparty/ffmpeg/mediacodec_surface.h

3rdparty/ffmpeg/mediacodec_wrapper.h:
	curl -o 3rdparty/ffmpeg/mediacodec_wrapper.h 'https://raw.githubusercontent.com/FFmpeg/FFmpeg/refs/heads/master/libavcodec/mediacodec_wrapper.h'

3rdparty/ffmpeg/mediacodec_surface.h:
	curl -o 3rdparty/ffmpeg/mediacodec_surface.h 'https://raw.githubusercontent.com/FFmpeg/FFmpeg/refs/heads/master/libavcodec/mediacodec_surface.h'

test:
	go build .
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows go build -tags crosscompile .
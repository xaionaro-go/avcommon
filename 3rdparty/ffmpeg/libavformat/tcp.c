#include <stdint.h>
#include <libavformat/avformat.h>

// copied from: https://github.com/FFmpeg/FFmpeg/blob/a0349d64fdd12295b826cd8db061922bdd050175/libavformat/tcp.c#L35-L50

typedef struct TCPContext {
    const AVClass *class;
    int fd;
    int listen;
    char *local_port;
    char *local_addr;
    int open_timeout;
    int rw_timeout;
    int listen_timeout;
    int recv_buffer_size;
    int send_buffer_size;
    int tcp_nodelay;
#if !HAVE_WINSOCK2_H
    int tcp_mss;
#endif /* !HAVE_WINSOCK2_H */
} TCPContext;

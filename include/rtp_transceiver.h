#ifndef EXAMPLES_RAWRTC_RTP_TRANSCEIVER_H_
#define EXAMPLES_RAWRTC_RTP_TRANSCEIVER_H_

#include <stdlib.h>

static const char* RTP_TRANSCEIVER_DIRECTION_SENDRECV = "sendrecv";
static const char* RTP_TRANSCEIVER_DIRECTION_SENDONLY = "sendonly";
static const char* RTP_TRANSCEIVER_DIRECTION_RECVONLY = "recvonly";
static const char* RTP_TRANSCEIVER_DIRECTION_INACTIVE = "inactive";
static const char* RTP_TRANSCEIVER_DIRECTION_STOPPED = "stopped";

typedef struct {
  char* direction;
  size_t size;
  char** stream_ids;
} raw_rtp_transceiver_init_t;

#endif  // EXAMPLES_RAWRTC_RTP_TRANSCEIVER_H_

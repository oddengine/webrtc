#ifndef EXAMPLES_RAWRTC_RTP_PARAMETERS_H_
#define EXAMPLES_RAWRTC_RTP_PARAMETERS_H_

#include <stdlib.h>

typedef struct {
  void* fd;
  char* mime_type;
  int clock_rate;
  int channels;
} raw_rtp_codec_capability_t;

typedef struct {
  raw_rtp_codec_capability_t* codecs;
  size_t size;
} raw_rtp_capabilities_t;

#endif  // EXAMPLES_RAWRTC_RTP_PARAMETERS_H_

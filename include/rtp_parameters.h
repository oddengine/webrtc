#ifndef EXAMPLES_RAWRTC_RTP_PARAMETERS_H_
#define EXAMPLES_RAWRTC_RTP_PARAMETERS_H_

#include <stdlib.h>

typedef struct {
  void* fd;
  char* mime_type;
  int clock_rate;
  int channels;
  char* sdp_fmtp_line;
} raw_rtp_codec_capability_t;

typedef struct {
  raw_rtp_codec_capability_t* codecs;
  size_t size;
} raw_rtp_capabilities_t;

typedef struct {
  void* fd;
  int payload_type;
  char* mime_type;
  int clock_rate;
  int channels;
  char* sdp_fmtp_line;
} raw_rtp_codec_parameters_t;

typedef struct {
  raw_rtp_codec_parameters_t* codecs;
  size_t size;
} raw_rtp_parameters_t;

#endif  // EXAMPLES_RAWRTC_RTP_PARAMETERS_H_

#ifndef EXAMPLES_RAWRTC_ICE_CANDIDATE_H_
#define EXAMPLES_RAWRTC_ICE_CANDIDATE_H_

typedef struct {
  char* candidate;
  char* sdp_mid;
  int sdp_mline_index;
} raw_ice_candidate_t;

#endif  // EXAMPLES_RAWRTC_ICE_CANDIDATE_H_

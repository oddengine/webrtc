#ifndef EXAMPLES_RAWRTC_API_H_
#define EXAMPLES_RAWRTC_API_H_

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#ifdef _WIN32
#define __dllexport__ __declspec(dllexport)
#else
#define __dllexport__
#endif

typedef struct {
  int64_t keyframe_interval;
} raw_rtc_constraints_t;

typedef struct {
  struct {
    int maxsize;
  } rotation;
} raw_default_writer_constraints_t;

typedef struct {
  void (*onresize)(void* target);
} raw_default_writer_observer_t;

typedef struct {
  void (*onsignalingchange)(void* target, const char* new_state);
  void (*ondatachannel)(void* target, void* data_channel);
  void (*onrenegotiationneeded)(void* target);
  void (*onconnectionchange)(void* target, const char* new_state);
  void (*oniceconnectionchange)(void* target, const char* new_state);
  void (*onicegatheringchange)(void* target, const char* new_state);
  void (*onicecandidate)(void* target,
                         const char* candidate,
                         const char* sdp_mid,
                         int sdp_mline_index);
  void (*onicecandidateerror)(void* target,
                              const char* address,
                              int port,
                              const char* url,
                              int error_code,
                              const char* error_text);
  void (*ontrack)(void* target, void* transceiver);
} raw_peer_connection_observer_t;

typedef struct {
  void (*onsuccess)(void* target, const char* type, const char* sdp);
  void (*onfailure)(void* target, const char* name, const char* message);
} raw_create_session_description_observer_t;

typedef struct {
  void (*onsuccess)(void* target);
  void (*onfailure)(void* target, const char* name, const char* message);
} raw_set_session_description_observer_t;

inline void* raw_calloc(size_t size) {
  void* p = malloc(size);
  if (p) {
    memset(p, 0, size);
  }
  return p;
}

inline void raw_free(void* p) {
  free(p);
}

#endif  // EXAMPLES_RAWRTC_API_H_

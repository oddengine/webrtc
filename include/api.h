#ifndef EXAMPLES_RAWRTC_API_H_
#define EXAMPLES_RAWRTC_API_H_

#include <stdlib.h>

typedef struct {
  size_t size;
  void** elements;
} raw_array_t;

typedef struct {
  void (*trace)(const char* message);
  void (*debug)(int n, const char* message);
  void (*info)(const char* message);
  void (*warn)(const char* message);
  void (*error)(const char* message);
} raw_logger_t;

typedef struct {
  void (*onsignalingchange)(void* observer, const char* new_state);
  void (*ondatachannel)(void* observer, void* data_channel);
  void (*onrenegotiationneeded)(void* observer);
  void (*onconnectionchange)(void* observer, const char* new_state);
  void (*oniceconnectionchange)(void* observer, const char* new_state);
  void (*onicegatheringchange)(void* observer, const char* new_state);
  void (*onicecandidate)(void* observer,
                         const char* candidate,
                         const char* sdp_mid,
                         int sdp_mline_index);
  void (*onicecandidateerror)(void* observer,
                              const char* address,
                              int port,
                              const char* url,
                              int error_code,
                              const char* error_text);
  void (*ontrack)(void* observer, void* transceiver);
} raw_peer_connection_observer_t;

typedef struct {
  void (*onsuccess)(void* observer, const char* type, const char* sdp);
  void (*onfailure)(void* observer, const char* name, const char* message);
} raw_create_session_description_observer_t;

typedef struct {
  void (*onsuccess)(void* observer);
  void (*onfailure)(void* observer, const char* name, const char* message);
} raw_set_session_description_observer_t;

#endif  // EXAMPLES_RAWRTC_API_H_

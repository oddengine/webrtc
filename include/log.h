#ifndef EXAMPLES_RAWRTC_LOG_LOG_H_
#define EXAMPLES_RAWRTC_LOG_LOG_H_

#include <stdlib.h>

typedef struct {
  struct {
    int maxsize;
  } rotation;
} raw_default_writer_constraints_t;

typedef struct {
  void (*onresize)();
} raw_default_writer_observer_t;

#endif  // EXAMPLES_RAWRTC_LOG_LOG_H_

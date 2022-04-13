#ifndef RAWRTC_LOG_LOG_H_
#define RAWRTC_LOG_LOG_H_

#include <string.h>

#include "../include/log.h"

typedef void *(*__create_default_logger_factory_fptr__)(void *fd, void *out, int level);
typedef void *(*__create_default_logger_fptr__)(void *fd, void *factory, const char *scope);
typedef void *(*__create_default_writer_fptr__)(void *fd, raw_default_writer_constraints_t *constraints);

typedef int (*__writer_open_fptr__)(void *writer, const char *path);
typedef int (*__writer_write_fptr__)(void *writer, const char *message, size_t size);
typedef int (*__writer_close_fptr__)(void *writer);

__create_default_logger_factory_fptr__ __create_default_logger_factory__;
__create_default_logger_fptr__ __create_default_logger__;
__create_default_writer_fptr__ __create_default_writer__;

__writer_open_fptr__ __writer_open__;
__writer_write_fptr__ __writer_write__;
__writer_close_fptr__ __writer_close__;

void *CreateDefaultLoggerFactory(void *fd, void *out, int level);
void *CreateDefaultLogger(void *fd, void *factory, const char *scope);
void *CreateDefaultWriter(void *fd, raw_default_writer_constraints_t *constraints);

int WriterOpen(void *writer, const char *path);
int WriterWrite(void *writer, const char *message, size_t size);
int WriterClose(void *writer);

#endif // RAWRTC_LOG_LOG_H_

#include "log.h"

void *CreateDefaultLoggerFactory(void *fd, void *out, int level)
{
    // __debugf__(6, "===> CreateDefaultLoggerFactory()");
    return __create_default_logger_factory__(fd, out, level);
}

void *CreateDefaultLogger(void *fd, void *factory, const char *scope)
{
    // __debugf__(6, "===> CreateDefaultLogger()");
    return __create_default_logger__(fd, factory, scope);
}

void *CreateDefaultWriter(void *fd, raw_default_writer_constraints_t *constraints)
{
    // __debugf__(6, "===> CreateDefaultWriter()");
    return __create_default_writer__(fd, constraints);
}

int WriterOpen(void *writer, const char *path)
{
    // __debugf__(6, "===> WriterOpen()");
    return __writer_open__(writer, path);
}

int WriterWrite(void *writer, const char *message, size_t size)
{
    // __debugf__(6, "===> WriterWrite()");
    return __writer_write__(writer, message, size);
}

int WriterClose(void *writer)
{
    // __debugf__(6, "===> WriterClose()");
    return __writer_close__(writer);
}

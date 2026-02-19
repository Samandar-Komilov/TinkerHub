#include "include/debug.h"
#include <errno.h>
#include <stdio.h>
#include <string.h>

static const char *debug_level_to_text(DebugLevel level) {
    switch (level) {
    case DEBUG_LEVEL_INFO:
        return "INFO";
    case DEBUG_LEVEL_WARN:
        return "WARN";
    case DEBUG_LEVEL_ERROR:
        return "ERROR";
    default:
        return "UNKNOWN";
    }
}

static void stderr_log(DebugLevel level, const char *component,
                       const char *message) {
    fprintf(stderr, "[%s] %s: %s\n", debug_level_to_text(level),
            component ? component : "app", message ? message : "(null)");
}

void debug_log(const DebugSink *sink, DebugLevel level, const char *component,
               const char *message) {
    if (!sink || !sink->log) {
        return;
    }
    sink->log(level, component, message);
}

void debug_log_errno(const DebugSink *sink, const char *component,
                     const char *operation) {
    char buf[256];
    int err = errno;
    snprintf(buf, sizeof(buf), "%s failed: errno=%d (%s)",
             operation ? operation : "operation", err, strerror(err));
    debug_log(sink, DEBUG_LEVEL_ERROR, component, buf);
}

const DebugSink STDERR_DEBUG_SINK = { .log = stderr_log };

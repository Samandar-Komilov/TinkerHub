#ifndef DEBUG_H
#define DEBUG_H

typedef enum DebugLevel {
    DEBUG_LEVEL_INFO = 0,
    DEBUG_LEVEL_WARN = 1,
    DEBUG_LEVEL_ERROR = 2,
} DebugLevel;

typedef struct DebugSink {
    void (*log)(DebugLevel level, const char *component, const char *message);
} DebugSink;

void debug_log(const DebugSink *sink, DebugLevel level, const char *component,
               const char *message);
void debug_log_errno(const DebugSink *sink, const char *component,
                     const char *operation);

extern const DebugSink STDERR_DEBUG_SINK;

#endif // DEBUG_H

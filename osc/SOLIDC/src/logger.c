#include "include/logger.h"
#include <errno.h>

int log_transaction(const Logger *lg, const Transaction *t,
                    const DebugSink *debug_sink) {
    char buf[MAX_BUFFER_SIZE];
    int n;

    if (!lg || !t || !lg->formatter || !lg->formatter->format || !lg->sender ||
        !lg->sender->send) {
        debug_log(debug_sink, DEBUG_LEVEL_ERROR, "logger",
                  "invalid logger dependencies");
        return -1;
    }

    n = lg->formatter->format(t, buf, MAX_BUFFER_SIZE);
    if (n < 0 || n >= MAX_BUFFER_SIZE) {
        debug_log(debug_sink, DEBUG_LEVEL_ERROR, "logger",
                  "formatter failed or produced invalid length");
        return -1;
    }

    errno = 0;
    if (lg->sender->send(buf, (size_t)n) < 0) {
        if (errno != 0) {
            debug_log_errno(debug_sink, "logger", "sender->send");
        } else {
            debug_log(debug_sink, DEBUG_LEVEL_ERROR, "logger",
                      "sender->send failed");
        }
        return -1;
    }

    return 0;
}

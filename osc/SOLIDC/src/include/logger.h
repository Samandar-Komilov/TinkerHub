#ifndef LOGGER_H
#define LOGGER_H

#include "config.h"
#include "debug.h"
#include "interfaces.h"

typedef struct Logger {
    const Formatter *formatter;
    const Sender *sender;
} Logger;

int log_transaction(const Logger *lg, const Transaction *t,
                    const DebugSink *debug_sink);

#endif // LOGGER_H

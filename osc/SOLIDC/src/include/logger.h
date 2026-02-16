#ifndef LOGGER_H
#define LOGGER_H

#include "config.h"
#include "interfaces.h"

typedef struct Logger {
    const Formatter *formatter;
    const Transport *transport;
} Logger;

int log_transaction(const Logger *lg, const Transaction *t);

#endif // LOGGER_H
#include "include/logger.h"

int log_transaction(const Logger *lg, const Transaction *t) {
    char buf[MAX_BUFFER_SIZE];
    int n = lg->formatter->format(t, buf, MAX_BUFFER_SIZE);
    if (n < 0) return -1;
    return lg->transport->send(buf, (size_t)n);
}
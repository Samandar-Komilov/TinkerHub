#ifndef INTERFACES_H
#define INTERFACES_H

#include "dto.h"
#include <stdlib.h>

typedef struct Formatter {
    /*
     * Contract:
     * - return bytes written in range [0, buf_size)
     * - return -1 on error
     * - must not write past buf_size
     */
    int (*format)(const Transaction *t, char *buf, size_t buf_size);
} Formatter;

typedef struct Sender {
    /* Contract: return 0 on success, -1 on error. */
    int (*send)(const char *msg, size_t len);
} Sender;

typedef struct Flushable {
    /* Optional capability: flush buffered data if any. */
    int (*flush)(void);
} Flushable;

typedef struct Connectable {
    /* Optional capability: connect/disconnect lifecycle for transport. */
    int (*connect)(void);
    int (*disconnect)(void);
} Connectable;

typedef struct Transport {
    const Sender *sender;
    const Flushable *flushable;
    const Connectable *connectable;
} Transport;

#endif // INTERFACES_H

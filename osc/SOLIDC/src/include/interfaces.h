#ifndef INTERFACES_H
#define INTERFACES_H

#include "dto.h"
#include <stdlib.h>

typedef struct Formatter {
    int (*format)(const Transaction* t, char *buf, size_t buf_size);
} Formatter;

typedef struct Transport {
    int (*send)(const char *msg, size_t len);
} Transport;

#endif // INTERFACES_H
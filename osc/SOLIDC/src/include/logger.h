#ifndef LOGGER_H
#define LOGGER_H

#include "config.h"
#include "dto.h"

void log_transaction(char *buffer, const Transaction *t,
                     const enum FormatTypeEnum ftype,
                     const enum TransportTypeEnum ttype);

#endif // LOGGER_H
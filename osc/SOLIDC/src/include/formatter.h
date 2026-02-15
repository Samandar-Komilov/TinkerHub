#ifndef FORMATTER_H
#define FORMATTER_H

#include "config.h"
#include "dto.h"

void format_transaction(const Transaction *t, const enum FormatTypeEnum ftype,
                        char *buf);

#endif // FORMATTER_H
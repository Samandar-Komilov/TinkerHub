#include "include/logger.h"
#include "include/dto.h"
#include "include/formatter.h"
#include "include/transport.h"

void log_transaction(char *buffer, const Transaction *t,
                     const enum FormatTypeEnum ftype,
                     const enum TransportTypeEnum ttype) {
    format_transaction(t, ftype, buffer);
    transport_log_message(ttype, buffer);
}
#ifndef TRANSPORT_H
#define TRANSPORT_H

#include "config.h"
#include "dto.h"

void transport_log_message(const enum TransportTypeEnum ttype, char *buffer);

#endif // TRANSPORT_H
#ifndef COMPONENTS_H
#define COMPONENTS_H

#include "interfaces.h"

extern const Formatter TEXT_FORMATTER;
extern const Formatter JSON_FORMATTER;

extern const Sender DISK_SENDER;
extern const Sender TCP_SENDER;
extern const Sender UDP_SENDER;

extern const Flushable DISK_FLUSHABLE;
extern const Connectable TCP_CONNECTABLE;

extern const Transport DISK_TRANSPORT;
extern const Transport TCP_TRANSPORT;
extern const Transport UDP_TRANSPORT;

#endif // COMPONENTS_H

#ifndef CONTROLLER_H
#define CONTROLLER_H

#include "dto.h"

int should_log_on_network(const Transaction *t);
void process_transaction(const Transaction *t);

#endif // CONTROLLER_H
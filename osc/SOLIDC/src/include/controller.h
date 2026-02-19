#ifndef CONTROLLER_H
#define CONTROLLER_H

#include "dto.h"
#include "debug.h"
#include "interfaces.h"
#include "logger.h"

typedef int (*LogPolicy)(const Transaction *t);

typedef struct AppContext {
    Logger local_logger;
    Logger network_logger;
    LogPolicy should_log_on_network;
    const Flushable *local_flushable;
    const Connectable *network_connectable;
    const DebugSink *debug_sink;
} AppContext;

int should_log_on_network(const Transaction *t);
int app_context_start(const AppContext *ctx);
int process_transaction_with_ctx(const AppContext *ctx, const Transaction *t);
int app_context_stop(const AppContext *ctx);

#endif // CONTROLLER_H

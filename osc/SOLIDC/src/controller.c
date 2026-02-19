#include "include/controller.h"
#include "include/config.h"

int should_log_on_network(const Transaction *t) {
    // Policy to WHEN to log
    if (!t) {
        return 0;
    }
    return t->amount <= MAX_TRANSACTION_AMOUNT_TO_LOG;
}

int app_context_start(const AppContext *ctx) {
    if (!ctx) {
        return -1;
    }
    if (ctx->network_connectable && ctx->network_connectable->connect) {
        if (ctx->network_connectable->connect() < 0) {
            debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                      "network connect capability failed");
            return -1;
        }
    }
    return 0;
}

int process_transaction_with_ctx(const AppContext *ctx, const Transaction *t) {
    if (!ctx || !t || !ctx->should_log_on_network) {
        if (ctx) {
            debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                      "invalid input or missing policy");
        }
        return -1;
    }

    if (log_transaction(&ctx->local_logger, t, ctx->debug_sink) < 0) {
        debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                  "local logging failed");
        return -1;
    }

    if (ctx->should_log_on_network(t)) {
        if (log_transaction(&ctx->network_logger, t, ctx->debug_sink) < 0) {
            debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                      "network logging failed");
            return -1;
        }
    }

    return 0;
}

int app_context_stop(const AppContext *ctx) {
    if (!ctx) {
        return -1;
    }

    if (ctx->local_flushable && ctx->local_flushable->flush &&
        ctx->local_flushable->flush() < 0) {
        debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                  "local flush capability failed");
        return -1;
    }
    if (ctx->network_connectable && ctx->network_connectable->disconnect &&
        ctx->network_connectable->disconnect() < 0) {
        debug_log(ctx->debug_sink, DEBUG_LEVEL_ERROR, "controller",
                  "network disconnect capability failed");
        return -1;
    }

    return 0;
}

#include "controller.h"
#include "include/components.h"
#include <stdio.h>
#include <stdlib.h>

int main() {
    AppContext ctx = {
        .local_logger =
            {
                .formatter = &TEXT_FORMATTER,
                .sender = DISK_TRANSPORT.sender,
            },
        .network_logger =
            {
                .formatter = &JSON_FORMATTER,
                .sender = TCP_TRANSPORT.sender,
            },
        .should_log_on_network = should_log_on_network,
        .local_flushable = DISK_TRANSPORT.flushable,
        .network_connectable = TCP_TRANSPORT.connectable,
        .debug_sink = &STDERR_DEBUG_SINK,
    };

    int stop = 0;
    unsigned int tid = 0;
    char user[20];
    double amount = 0.0;

    if (app_context_start(&ctx) < 0) {
        debug_log(ctx.debug_sink, DEBUG_LEVEL_WARN, "main",
                  "application context start failed; continuing");
    }

    while (!stop) {
        printf("Enter transaction (tid user amount): ");
        if (scanf("%u %19s %lf", &tid, user, &amount) != 3) {
            fprintf(stderr, "invalid input\n");
            return EXIT_FAILURE;
        }

        Transaction t = {tid, user, amount};
        if (process_transaction_with_ctx(&ctx, &t) < 0) {
            debug_log(ctx.debug_sink, DEBUG_LEVEL_ERROR, "main",
                      "failed to process transaction");
        }

        printf("Stop? (0/1): ");
        if (scanf("%d", &stop) != 1) {
            fprintf(stderr, "invalid stop value\n");
            return EXIT_FAILURE;
        }
    }

    if (app_context_stop(&ctx) < 0) {
        debug_log(ctx.debug_sink, DEBUG_LEVEL_ERROR, "main",
                  "failed to finalize application context");
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}

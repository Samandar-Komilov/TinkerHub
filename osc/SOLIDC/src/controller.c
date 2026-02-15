#include "include/controller.h"
#include "include/config.h"
#include "include/logger.h"

int should_log_on_network(const Transaction *t) {
    // Policy to WHEN to log
    return t->amount <= MAX_TRANSACTION_AMOUNT_TO_LOG;
}

void process_transaction(const Transaction *t) {
    char buf[MAX_BUFFER_SIZE];
    log_transaction(buf, t, TEXT, DISK); // mechanism orchestration
    if (should_log_on_network(t)) {
        log_transaction(buf, t, JSON, TCP);
        // log_transaction(buf, t, JSON, UDP);
    }
}
#include "include/controller.h"
#include "include/config.h"
#include "include/components.h"
#include "include/logger.h"

static const Logger DISK_TEXT_LOGGER = {
    .formatter = &TEXT_FORMATTER,
    .transport = &DISK_TRANSPORT,
};

static const Logger NETWORK_JSON_LOGGER = {
    .formatter = &JSON_FORMATTER,
    .transport = &TCP_TRANSPORT,
};

int should_log_on_network(const Transaction *t) {
    // Policy to WHEN to log
    return t->amount <= MAX_TRANSACTION_AMOUNT_TO_LOG;
}

void process_transaction(const Transaction *t) {
    (void)log_transaction(&DISK_TEXT_LOGGER, t);

    if (should_log_on_network(t)) {
        (void)log_transaction(&NETWORK_JSON_LOGGER, t);
    }
}

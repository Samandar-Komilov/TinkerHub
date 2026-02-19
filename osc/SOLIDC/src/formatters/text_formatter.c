#include "interfaces.h"
#include <stdio.h>

static int text_format(const Transaction *t, char *out, size_t sz) {
    int n;

    if (!t || !out || sz == 0 || !t->user) {
        return -1;
    }

    n = snprintf(out, sz, "Transaction %u: User %s sent %.2f\n", t->tid,
                 t->user, t->amount);
    if (n < 0 || (size_t)n >= sz) {
        return -1;
    }

    return n;
}
const Formatter TEXT_FORMATTER = { .format = text_format };

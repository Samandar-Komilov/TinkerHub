#include "interfaces.h"
#include <stdio.h>

static int text_format(const Transaction *t, char *out, size_t sz) {
    return snprintf(out, sz, "Transaction %u: User %s sent %.2f\n",
                    t->tid, t->user, t->amount);
}
const Formatter TEXT_FORMATTER = { .format = text_format };

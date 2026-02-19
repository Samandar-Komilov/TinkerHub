#include "interfaces.h"
#include "config.h"
#include "stdio.h"

static int disk_send(const char *msg, size_t len) {
    int rc = -1;
    FILE *fp = fopen(LOG_FILE, "a");
    size_t w;

    if (!msg) {
        return -1;
    }
    if (!fp) {
        return -1;
    }

    w = fwrite(msg, 1, len, fp);
    if (w == len) {
        rc = 0;
    }

    fclose(fp);
    return rc;
}

static int disk_flush(void) {
    return 0;
}

const Sender DISK_SENDER = { .send = disk_send };
const Flushable DISK_FLUSHABLE = { .flush = disk_flush };
const Transport DISK_TRANSPORT = {
    .sender = &DISK_SENDER,
    .flushable = &DISK_FLUSHABLE,
    .connectable = NULL,
};

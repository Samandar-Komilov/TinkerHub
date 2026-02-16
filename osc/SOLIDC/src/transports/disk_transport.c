#include "interfaces.h"
#include "config.h"
#include "stdio.h"

static int disk_send(const char *msg, size_t len) {
    FILE *fp = fopen(LOG_FILE, "a");
    if (!fp) return -1;
    size_t w = fwrite(msg, 1, len, fp);
    fclose(fp);
    return (w == len) ? 0 : -1;
}
const Transport DISK_TRANSPORT = { .send = disk_send };

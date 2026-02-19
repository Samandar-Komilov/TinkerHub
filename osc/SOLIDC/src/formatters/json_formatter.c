#include "interfaces.h"
#include "cJSON.h"
#include <stdio.h>
#include <stdlib.h>

static int json_format(const Transaction *t, char *out, size_t sz) {
    int n;
    cJSON *root = cJSON_CreateObject();
    char *json_str;

    if (!t || !out || sz == 0 || !t->user || !root) {
        cJSON_Delete(root);
        return -1;
    }

    if (!cJSON_AddNumberToObject(root, "tid", t->tid) ||
        !cJSON_AddStringToObject(root, "user", t->user) ||
        !cJSON_AddNumberToObject(root, "amount", t->amount)) {
        cJSON_Delete(root);
        return -1;
    }

    json_str = cJSON_PrintUnformatted(root);
    if (!json_str) {
        cJSON_Delete(root);
        return -1;
    }

    n = snprintf(out, sz, "%s", json_str);
    free(json_str);
    cJSON_Delete(root);

    if (n < 0 || (size_t)n >= sz) {
        return -1;
    }

    return n;
}

const Formatter JSON_FORMATTER = { .format = json_format };

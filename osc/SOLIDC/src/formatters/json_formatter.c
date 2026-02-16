#include "interfaces.h"
#include "cJSON.h"
#include <stdio.h>
#include <stdlib.h>

static int json_format(const Transaction *t, char *out, size_t sz) {
    cJSON *root = cJSON_CreateObject();
    cJSON_AddNumberToObject(root, "tid", t->tid);
    cJSON_AddStringToObject(root, "user", t->user);
    cJSON_AddNumberToObject(root, "amount", t->amount);
    char *json_str = cJSON_PrintUnformatted(root);
    int n = snprintf(out, sz, "%s", json_str);
    free(json_str);
    cJSON_Delete(root);

    return n;
}

const Formatter JSON_FORMATTER = { .format = json_format };

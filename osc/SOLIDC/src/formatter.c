#include "include/formatter.h"
#include "json/cJSON.h"
#include <stdio.h>
#include <stdlib.h>

void format_transaction(const Transaction *t, const enum FormatTypeEnum ftype,
                        char *buf) {
    switch (ftype) {
    case TEXT: {
        snprintf(buf, MAX_BUFFER_SIZE, "Transaction %d: User %s sent amount %.2f", t->tid,
                t->user, t->amount);
        break;
    }
    case JSON: {
        cJSON *root = cJSON_CreateObject();
        cJSON_AddNumberToObject(root, "tid", t->tid);
        cJSON_AddStringToObject(root, "user", t->user);
        cJSON_AddNumberToObject(root, "amount", t->amount);
        char *json_str = cJSON_PrintUnformatted(root);
        snprintf(buf, MAX_BUFFER_SIZE, "%s", json_str);
        free(json_str);
        cJSON_Delete(root);
    } break;
    }
}
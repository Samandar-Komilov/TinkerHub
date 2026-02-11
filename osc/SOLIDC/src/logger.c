#include "logger.h"
#include "config.h"
#include <stdio.h>

void log_transaction(const int tid, const char *user, const double amount) {
  FILE *fp = fopen(LOG_FILE, "a");
  if (fp == NULL) {
    return;
  }

  fprintf(fp, "Transaction %d: User %s sent amount %.2f\n", tid, user, amount);
  fclose(fp);
}
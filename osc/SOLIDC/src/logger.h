typedef struct {
  int tid;
  const char *user;
  double amount;
} Transaction;

void log_transaction(const Transaction *t);
#include "logger.h"
#include "json/cJSON.h"
#include <stdio.h>
#include <stdlib.h>

int main() {
  Transaction t = {1, "John", 10.0};
  log_transaction(&t);
  printf("Hello World\n");
}
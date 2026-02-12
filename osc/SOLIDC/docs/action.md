# Action | Step by Step Build Log


### F1: Core Logging

Initially, we are expected to create a minimal yet working logging system that prints logs to a specified file:

```c
// logger.h
void log_transaction(const int tid, const char *user, const double amount);
```

```c
// logger.c
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
```

```c
// main.c
#include "logger.h"

int main() {
  log_transaction(1, "John", 10.0);
  return 0;
}
```

Well, nothing special, nothing extraordinary. But recall that our requirements change frequently. 

### F2: Send Logs over a Network in JSON, but disk logging should still be text and work

We need to add JSON logging to a network endpoint. This means we need to add new functionality in the existing function, based on our current expertise. The worst scenario would be both JSON construction and network logic in the same function.

We define a `Transaction` struct to hold the transaction data.

```c
// logger.h
typedef struct {
  int tid;
  const char *user;
  double amount;
} Transaction;

void log_transaction(const Transaction *t);
```

A popular JSON library `cJSON` is used for easy JSON serialization and deserialization in C.

```c
// logger.c
#include "logger.h"
#include "config.h"
#include "json/cJSON.h"
#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

void log_transaction(const Transaction *t) {
  char msg[256];
  int msg_len = sprintf(msg, "Transaction %d: User %s sent amount %.2f\n",
                        t->tid, t->user, t->amount);

  // Disk log
  FILE *fp = fopen(LOG_FILE, "a");
  if (fp == NULL) {
    return;
  }
  fprintf(fp, "%s", msg);
  fclose(fp);

  // 1. Create Socket
  int sockfd = socket(AF_INET, SOCK_DGRAM, 0);
  if (sockfd < 0)
    return;

  // 2. Setup Address Structure
  struct sockaddr_in addr;
  memset(&addr, 0, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_port = htons(LOG_PORT);

  // 3. Convert IPv4 address from text to binary
  if (inet_pton(AF_INET, LOG_HOST, &addr.sin_addr) <= 0) {
    close(sockfd);
    return;
  }

  // 4. JSON construction
  cJSON *root = cJSON_CreateObject();
  cJSON_AddNumberToObject(root, "tid", t->tid);
  cJSON_AddStringToObject(root, "user", t->user);
  cJSON_AddNumberToObject(root, "amount", t->amount);
  char *json_str = cJSON_PrintUnformatted(root);

  // 5. Send Packet
  sendto(sockfd, json_str, strlen(json_str), 0, (struct sockaddr *)&addr,
         sizeof(addr));

  cJSON_Delete(root);
  close(sockfd);
}
```

Did you notice the amount of change needed? The function which was just writing to a file is now deciding what to do, and HOW to do. And this is exactly the violation **Single Responsibility Principle** says we should avoid.

##### What To Do?


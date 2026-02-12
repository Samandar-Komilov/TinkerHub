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

  FILE *fp = fopen(LOG_FILE, "a");
  if (fp == NULL) {
    return;
  }

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
#include "interfaces.h"
#include "config.h"
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

static int tcp_send(const char *msg, size_t len) {
    int sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0)
        return -1;

    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(LOG_PORT);
    if (inet_pton(AF_INET, LOG_HOST, &addr.sin_addr) <= 0) {
        close(sockfd);
        return -1;
    }

    if (connect(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        close(sockfd);
        return -1;
    }
    ssize_t sent = send(sockfd, msg, len, 0);
    close(sockfd);
    return sent >= 0 ? 0 : -1;
}

const Transport TCP_TRANSPORT = { .send = tcp_send };

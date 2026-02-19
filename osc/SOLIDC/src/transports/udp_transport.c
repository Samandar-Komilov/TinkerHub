#include "interfaces.h"
#include "config.h"
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

static int udp_send(const char *msg, size_t len) {
    ssize_t sent;
    int sockfd = socket(AF_INET, SOCK_DGRAM, 0);
    if (!msg) {
        return -1;
    }
    if (sockfd < 0) {
        return -1;
    }

    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(LOG_PORT);

    if (inet_pton(AF_INET, LOG_HOST, &addr.sin_addr) <= 0) {
        close(sockfd);
        return -1;
    }
    sent = sendto(sockfd, msg, len, 0, (struct sockaddr *)&addr, sizeof(addr));
    close(sockfd);
    return sent == (ssize_t)len ? 0 : -1;
}

const Sender UDP_SENDER = { .send = udp_send };
const Transport UDP_TRANSPORT = {
    .sender = &UDP_SENDER,
    .flushable = NULL,
    .connectable = NULL,
};

#include "interfaces.h"
#include "config.h"
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

static int udp_send(const char *msg, size_t len) {
    int sockfd = socket(AF_INET, SOCK_DGRAM, 0);
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
    ssize_t sent = sendto(sockfd, msg, len, 0, (struct sockaddr *)&addr,
                          sizeof(addr));
    close(sockfd);
    return sent >= 0 ? 0 : -1;
}

const Transport UDP_TRANSPORT = { .send = udp_send };

#include "interfaces.h"
#include "config.h"
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

static int tcp_open_socket(void) {
    int sockfd = socket(AF_INET, SOCK_STREAM, 0);
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

    if (connect(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        close(sockfd);
        return -1;
    }

    return sockfd;
}

static int tcp_send(const char *msg, size_t len) {
    int sockfd;
    size_t sent_total = 0;

    if (!msg) {
        return -1;
    }

    sockfd = tcp_open_socket();
    if (sockfd < 0) {
        return -1;
    }

    while (sent_total < len) {
        ssize_t sent =
            send(sockfd, msg + sent_total, len - sent_total, 0);
        if (sent <= 0) {
            close(sockfd);
            return -1;
        }
        sent_total += (size_t)sent;
    }

    if (close(sockfd) < 0) {
        return -1;
    }

    return 0;
}

static int tcp_connect_capability(void) {
    int sockfd = tcp_open_socket();
    if (sockfd < 0) {
        return -1;
    }
    close(sockfd);
    return 0;
}

static int tcp_disconnect_capability(void) {
    return 0;
}

const Sender TCP_SENDER = { .send = tcp_send };
const Connectable TCP_CONNECTABLE = {
    .connect = tcp_connect_capability,
    .disconnect = tcp_disconnect_capability,
};
const Transport TCP_TRANSPORT = {
    .sender = &TCP_SENDER,
    .flushable = NULL,
    .connectable = &TCP_CONNECTABLE,
};

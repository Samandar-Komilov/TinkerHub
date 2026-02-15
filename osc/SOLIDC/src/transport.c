#include "include/transport.h"
#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

void transport_log_message(const enum TransportTypeEnum ttype, char *buffer) {
    size_t len = strlen(buffer);
    
    if (len > 0 && buffer[len-1] != '\n' && len < MAX_BUFFER_SIZE - 1) {
        buffer[len] = '\n';
        buffer[len+1] = '\0';
        len++;
    }

    switch (ttype) {
    case DISK: {
        FILE *fp = fopen(LOG_FILE, "a");
        if (fp == NULL) {
            return;
        }
        fprintf(fp, "%s", buffer);
        fclose(fp);
        break;
    }
    case TCP: {
        int sockfd = socket(AF_INET, SOCK_STREAM, 0);
        if (sockfd < 0)
            return;

        struct sockaddr_in addr;
        memset(&addr, 0, sizeof(addr));
        addr.sin_family = AF_INET;
        addr.sin_port = htons(LOG_PORT);
        if (inet_pton(AF_INET, LOG_HOST, &addr.sin_addr) <= 0) {
            close(sockfd);
            return;
        }

        if (connect(sockfd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
            close(sockfd);
            return;
        }
        send(sockfd, buffer, strlen(buffer), 0);
        close(sockfd);
        break;
    }
    case UDP: {
        int sockfd = socket(AF_INET, SOCK_DGRAM, 0);
        if (sockfd < 0)
            return;

        struct sockaddr_in addr;
        memset(&addr, 0, sizeof(addr));
        addr.sin_family = AF_INET;
        addr.sin_port = htons(LOG_PORT);

        if (inet_pton(AF_INET, LOG_HOST, &addr.sin_addr) <= 0) {
            close(sockfd);
            return;
        }
        sendto(sockfd, buffer, strlen(buffer), 0, (struct sockaddr *)&addr,
               sizeof(addr));
        close(sockfd);
        break;
        }
    }
}
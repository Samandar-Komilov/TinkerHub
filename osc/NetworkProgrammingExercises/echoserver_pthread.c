#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <arpa/inet.h>
#include <sys/wait.h>
#include <signal.h>
#include <errno.h>

#define PORT "8010"
#define BACKLOG 5
#define BUFFER_SIZE 1024


void *handle_client_thread(void *arg);
void *get_in_addr(struct sockaddr *sa);

int main(int argc, char *argv[]){
    // Initializations (optional, we can directly declare and initialize actually)
    int sockfd;
    struct addrinfo hints, *servinfo, *p;
    struct sockaddr_storage their_addr;
    socklen_t sin_size;
    int yes = 1;
    char s[INET6_ADDRSTRLEN];

    memset(&hints, 0, sizeof(hints));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_flags = AI_PASSIVE;

    // Getting server address information using getaddrinfo() rather than manual
    int addr_got = getaddrinfo(NULL, PORT, &hints, &servinfo);
    if (addr_got != 0){
        fprintf(stderr, "[ERROR] getaddrinfo() has faced an error: %s\n", gai_strerror(addr_got));
        return 1;
    }

    // A single host can have multiple IP addresses, so we should iterate over and choose the first one that works
    // create socket descriptor, make it to use the same port again, and bind to the port.
    for (p=servinfo; p!=NULL; p=p->ai_next){
        if ((sockfd = socket(p->ai_family, p->ai_socktype, p->ai_protocol)) == -1){
            fprintf(stderr, "[ERROR] Error while opening a socket: %s\n", gai_strerror(addr_got));
            continue;
        }

        if (setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int)) == -1){
            perror("setsockopt");
            exit(1);
        }

        int is_binded = bind(sockfd, p->ai_addr, p->ai_addrlen);
        if (is_binded < 0){
            close(sockfd);
            fprintf(stderr, "[ERROR] Error while binding to a port: %s\n", gai_strerror(is_binded));
            continue;
        }

        break;
    }
    
    freeaddrinfo(servinfo);

    if (p == NULL){
        fprintf(stderr, "Server: Failed to bind.\n");
        return 1;
    }
    
    // Prepare the socket to start accepting connections. Tell the OS "this socket is now busy, waiting for connections
    // allocate queue of length 5 to handle multiple connections sequentially.
    int is_listened = listen(sockfd, 5);
    if (is_listened < 0){
        fprintf(stderr, "[ERROR] Error while listening to the socket: %s\n", gai_strerror(is_listened));
    }

    printf("===== WAITING FOR CONNECTIONS on port %s =====\n", PORT);

    while (1) {
        // Accept a connection once it comes, manage it with a new special socket descriptor
        sin_size = sizeof(their_addr);
        int *client_fd_ptr = (int*) malloc(sizeof(int));
        *client_fd_ptr = accept(sockfd, (struct sockaddr *)&their_addr, &sin_size);
        if (*client_fd_ptr == -1){
            perror("accept");
            continue;
        }
    
        inet_ntop(their_addr.ss_family, get_in_addr((struct sockaddr *)&their_addr), s, sizeof(s));
        printf("server: got connection from %s\n", s);

        // Create a new thread to handle the client. Detach it so that it works independently without blocking the main thread
        pthread_t thread_id;
        pthread_create(&thread_id, NULL, handle_client_thread, client_fd_ptr);
        pthread_detach(thread_id);
    }

    return 0;
}

void *handle_client_thread(void *arg){
    int client_fd = *(int*) arg;
    // Ownership moved from main thread to handler thread
    free(arg);
    
    char buff[BUFFER_SIZE];
    int bytes_received;

    while ((bytes_received = recv(client_fd, buff, BUFFER_SIZE, 0)) > 0){
        send(client_fd, buff, bytes_received, 0);
        printf("Received: %s\n", buff);
        memset(buff, 0, sizeof(buff));
    }

    close(client_fd);
    return NULL;
}

void *get_in_addr(struct sockaddr *sa){
    if (sa->sa_family == AF_INET){
        return &(((struct sockaddr_in*)sa)->sin_addr);
    }

    return &(((struct sockaddr_in6*)sa)->sin6_addr);
}
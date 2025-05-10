#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
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


void sigchld_handler(int s);
void *get_in_addr(struct sockaddr *sa);

int main(int argc, char *argv[]){
    // Initializations (optional, we can directly declare and initialize actually)
    int sockfd, new_fd;
    struct addrinfo hints, *servinfo, *p;
    struct sockaddr_storage their_addr;
    socklen_t sin_size;
    struct sigaction sa;
    int yes = 1;
    char s[INET6_ADDRSTRLEN];
    
    // "Registering" signal handler so that it will be triggered once child process becomes zombie
    sa.sa_handler = sigchld_handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = SA_RESTART;
    if (sigaction(SIGCHLD, &sa, NULL) == -1){
        perror("sigaction");
        exit(1);
    }

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
        new_fd = accept(sockfd, (struct sockaddr *)&their_addr, &sin_size);
        if (new_fd == -1){
            perror("accept");
            continue;
        }
    
        inet_ntop(their_addr.ss_family, get_in_addr((struct sockaddr *)&their_addr), s, sizeof(s));
        printf("server: got connection from %s\n", s);
    
        if (fork() == 0){
            // Once a request is accepted, fork a new process to handle it accordingly
            close(sockfd);
            char buffer[128];
            int bytes_received;
            
            while ((bytes_received = recv(new_fd, buffer, 128, 0)) > 0){
                // Receive the message in chunks, 128 bytes in each loop
                printf("Received: %s\n", buffer);
                send(new_fd, buffer, bytes_received, 0);
                memset(buffer, 0, 128);
            }
            close(new_fd);
            exit(0);
        }
        close(new_fd);
    }

    return 0;
}

void sigchld_handler(int s){
    // Reap all zombie processed (child processes which are terminated after doing their job)
    int saved_errno = errno;

    while (waitpid(-1, NULL, WNOHANG) > 0);

    errno = saved_errno;
}

void *get_in_addr(struct sockaddr *sa){
    if (sa->sa_family == AF_INET){
        return &(((struct sockaddr_in*)sa)->sin_addr);
    }

    return &(((struct sockaddr_in6*)sa)->sin6_addr);
}
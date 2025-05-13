#include <arpa/inet.h>
#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <netinet/in.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/epoll.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

#define PORT "8010"
#define BACKLOG 5
#define MAX_EVENTS 10
#define BUFFER_SIZE 1024

int make_socket_non_blocking(int fd);
void *get_in_addr(struct sockaddr *sa);

int main(int argc, char *argv[]) {
  // Initializations (optional, we can directly declare and initialize actually)
  int sockfd, epoll_fd;
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
  if (addr_got != 0) {
    fprintf(stderr, "[ERROR] getaddrinfo() has faced an error: %s\n",
            gai_strerror(addr_got));
    return 1;
  }

  // A single host can have multiple IP addresses, so we should iterate over and
  // choose the first one that works create socket descriptor, make it to use
  // the same port again, and bind to the port.
  for (p = servinfo; p != NULL; p = p->ai_next) {
    if ((sockfd = socket(p->ai_family, p->ai_socktype, p->ai_protocol)) == -1) {
      fprintf(stderr, "[ERROR] Error while opening a socket: %s\n",
              gai_strerror(addr_got));
      continue;
    }

    if (setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int)) == -1) {
      perror("setsockopt");
      exit(1);
    }

    int is_binded = bind(sockfd, p->ai_addr, p->ai_addrlen);
    if (is_binded < 0) {
      close(sockfd);
      fprintf(stderr, "[ERROR] Error while binding to a port: %s\n",
              gai_strerror(is_binded));
      continue;
    }

    break;
  }

  freeaddrinfo(servinfo);

  if (p == NULL) {
    fprintf(stderr, "Server: Failed to bind.\n");
    return 1;
  }

  // Prepare the socket to start accepting connections. Tell the OS "this socket
  // is now busy, waiting for connections allocate queue of length 5 to handle
  // multiple connections sequentially.
  int is_listened = listen(sockfd, 5);
  if (is_listened < 0) {
    fprintf(stderr, "[ERROR] Error while listening to the socket: %s\n",
            gai_strerror(is_listened));
  }

  // Use epoll() to handle multiple connections concurrently
  epoll_fd = epoll_create1(0);
  if (epoll_fd == -1) {
    fprintf(stderr, "[ERROR] epoll_create1() call");
    exit(1);
  }

  struct epoll_event ev, events[MAX_EVENTS];
  ev.events = EPOLLIN;
  ev.data.fd = sockfd;
  if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, sockfd, &ev) == -1) {
    perror("[ERROR] epoll_ctl: ");
    exit(1);
  }

  printf("===== WAITING FOR CONNECTIONS on port %s =====\n", PORT);

  while (1) {
    int n_ready = epoll_wait(epoll_fd, events, MAX_EVENTS, -1);
    if (n_ready == -1) {
      perror("[ERROR] epoll_wait");
      continue;
    }

    for (int i = 0; i < n_ready; i++) {
      if (events[i].data.fd == sockfd) {
        // Accept a new connection
        int client_fd =
            accept(sockfd, (struct sockaddr *)&their_addr, &sin_size);
        if (client_fd == -1) {
          perror("[ERROR] Error while accepting a new connection");
          continue;
        }

        make_socket_non_blocking(client_fd);

        ev.events = EPOLLIN | EPOLLOUT;
        ev.data.fd = client_fd;
        if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, client_fd, &ev) == -1) {
          perror("[ERROR] Error while responding to the client socket: ");
          continue;
        }

        printf("[INFO] Client FD: %d, Assigned FD: %d\n", client_fd,
               ev.data.fd);

        inet_ntop(their_addr.ss_family,
                  get_in_addr((struct sockaddr *)&their_addr), s, sizeof(s));
        printf("[INFO] Got connection from %s\n", s);
      } else {
        // Read from client socket and handle accordingly
        char buff[BUFFER_SIZE];
        int bytes_read = recv(events[i].data.fd, buff, BUFFER_SIZE, 0);
        if (bytes_read <= 0) {
          if (errno == EAGAIN || errno == EWOULDBLOCK) {
            // Nothing to read now — try again later
            continue;
          }
          if (bytes_read < 0)
            perror("[ERROR] Error while reading from client socket: ");
          close(events[i].data.fd);
          printf("[INFO] Connection closed.\n");
        } else {
          printf("[INFO] Received: %s\nBytes read: %d\n", buff);
          send(events[i].data.fd, buff, bytes_read, 0);
        }
      }
    }
  }

  close(epoll_fd);
  close(sockfd);

  return 0;
}

int make_socket_non_blocking(int fd) {
  int flags = fcntl(fd, F_GETFL, 0);
  if (flags == -1)
    return -1;
  return fcntl(fd, F_SETFL, flags | O_NONBLOCK);
}

void *get_in_addr(struct sockaddr *sa) {
  if (sa->sa_family == AF_INET) {
    return &(((struct sockaddr_in *)sa)->sin_addr);
  }

  return &(((struct sockaddr_in6 *)sa)->sin6_addr);
}
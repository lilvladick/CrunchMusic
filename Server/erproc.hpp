#ifndef ERPROC_H
#define ERPROC_H

#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

int Socket(int domain, int type, int protocol);

void Bind(int fd, const struct sockaddr *addr, socklen_t addrlen);

void Listen(int fd, int backlog);

int Accept(int fd, struct sockaddr *addr, socklen_t *addrLen);

#endif
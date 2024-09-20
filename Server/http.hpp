#ifndef HTTP_H
#define HTTP_H

#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

void handleConnection(int clientSocket);

#endif
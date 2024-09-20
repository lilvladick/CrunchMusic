#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

int Socket(int domain, int type, int protocol){
    int serversocket = socket(domain, type, protocol);
    if (serversocket < 0) {
        std::cout << "Error creating socket" << std::endl;
        return 1;
    }

    return serversocket;
}

void Bind(int socket, const sockaddr *addr, socklen_t addrLen){
    if (bind(socket, addr, addrLen) < 0) {
        std::cout << "Error binding socket" << std::endl;
        return 1;
    }
}

void Listen(int socket, int backlog){
    if (listen(socket, backlog) < 0) {
        std::cout << "Error listening on socket" << std::endl;
        return 1;
    }
}

int Accept(int socket,const sockaddr *addr, socklen_t addrLen ){
    int clientSocket = accept(socket, addr, addrLen);
    if (clientSocket < 0) {
        std::cout << "Error accepting connection" << std::endl;
        return 1;
    }
    
    return clientSocket;
}
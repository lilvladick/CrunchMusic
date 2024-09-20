#include "erproc.hpp"

int Socket(int domain, int type, int protocol){
    int serversocket = socket(domain, type, protocol);
    if (serversocket < 0) {
        std::cout << "Error creating socket" << std::endl;
        return 1;
    }

    return serversocket;
}

void Bind(int socket,const struct sockaddr *addr, socklen_t addrLen){
    if (bind(socket, addr, addrLen) < 0) {
        std::cout << "Error binding socket" << std::endl;
        return ;
    }
}

void Listen(int socket, int backlog){
    if (listen(socket, backlog) < 0) {
        std::cout << "Error listening on socket" << std::endl;
        return ;
    }
}

int Accept(int socket,struct sockaddr *addr, socklen_t *addrLen ){
    int clientSocket = accept(socket, addr, addrLen);
    if (clientSocket < 0) {
        std::cout << "Error accepting connection" << std::endl;
        return 1;
    }
    
    return clientSocket;
}
#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>
#include "erproc.hpp"
#include "http.hpp"
#include "postgre.hpp"
// Все коментарии и текст в ошибках написаны GigaCode.
// Так что рука нейросетей только в текст ошибок, ну почти...

#define PORT 8080

int main() {
    int serverSocket, clientSocket;
    struct sockaddr_in serverAddr, clientAddr;
    socklen_t clientAddrLen = sizeof(clientAddr);

    serverSocket = Socket(AF_INET, SOCK_STREAM, 0);
    
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(PORT);
    serverAddr.sin_addr.s_addr = INADDR_ANY;


    Bind(serverSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr));

    Listen(serverSocket,5);

    std::cout << "Waiting for connection..." << std::endl;

    while (true)
    {
        clientSocket = Accept(serverSocket, (struct sockaddr*)&clientAddr, &clientAddrLen);
        
        std::cout << "Connection accepted" << std::endl;

        handleConnection(clientSocket);
    }
}
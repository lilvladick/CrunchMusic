#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

// Все коментарии и текст в ошибках написаны GigaCode.
// Так что рука нейросетей только в текст ошибок, ну почти...

#define PORT 8080
#define BUFFER_SIZE 1024

void handleConnection(int clientSocket) {
    char buffer[BUFFER_SIZE];
    int bytesReceived = recv(clientSocket, buffer, BUFFER_SIZE, 0);
    if (bytesReceived <= 0) {
        close(clientSocket);
        return;
    }

    buffer[bytesReceived] = '\0';
    std::cout << "Received: " << buffer << std::endl;
    
    if (strstr(buffer, "GET / HTTP/1.1") != NULL) {
        const char* response = "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\nГО СОСЕТ У ПЛЮСОВ";
        send(clientSocket, response, strlen(response), 0);
    } else {
        const char* response = "HTTP/1.1 404 Not Found\r\nContent-Type: text/html\r\n\r\nДУРАК ВСЕ СЛОМАЛОСЬ";
        send(clientSocket, response, strlen(response), 0);
    }

    close(clientSocket);
}

int main() {
    int serverSocket, clientSocket;
    struct sockaddr_in serverAddr, clientAddr;
    socklen_t clientAddrLen = sizeof(clientAddr);

    serverSocket = socket(AF_INET, SOCK_STREAM, 0);

    if (serverSocket < 0) {
        std::cout << "Error creating socket" << std::endl;
        return 1;
    }

    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(PORT);
    serverAddr.sin_addr.s_addr = INADDR_ANY;

    if (bind(serverSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr)) < 0) {
        std::cout << "Error binding socket" << std::endl;
        return 1;
    }

    if (listen(serverSocket, 5) < 0) {
        std::cout << "Error listening on socket" << std::endl;
        return 1;
    }

    std::cout << "Waiting for connection..." << std::endl;

    while (true)
    {
        clientSocket = accept(serverSocket, (struct sockaddr*)&clientAddr, &clientAddrLen);
        if (clientSocket < 0) {
            std::cout << "Error accepting connection" << std::endl;
            return 1;
        }
        std::cout << "Connection accepted" << std::endl;

        handleConnection(clientSocket);
    }
}
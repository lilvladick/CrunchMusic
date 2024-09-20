#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

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
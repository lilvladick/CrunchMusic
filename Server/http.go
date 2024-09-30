package main

import (
	"fmt"
	"log"
	"strings"
	"syscall"
)

// Функция обработки запроса
// syscall.Handle - файловый дескриптор (syscall.Handle для винды, а int для unix)
func handleRequest(fd int) {
	defer syscall.Close(fd)
	buf := make([]byte, 1024)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		log.Print("read error: ", err)
		return
	}

	request := string(buf[:n])
	log.Print(request)

	if strings.HasPrefix(request, "GET ") {
		if strings.Contains(request, "/tracks") {
			query := "SELECT * FROM tracks"

			jsonData, err := getResultsJson(query)
			if err != nil {
				log.Print("error fetching tracks: ", err)
				errorResponse := "HTTP/1.1 500 Internal Server Error\r\n" +
					"Content-Type: text/plain\r\n" +
					"\r\n" +
					"Error fetching tracks"
				syscall.Write(fd, []byte(errorResponse))
				return
			}

			response := "HTTP/1.1 200 OK\r\n" +
				"Content-Type: application/json\r\n" +
				"Content-Length: " + fmt.Sprintf("%d", len(jsonData)) + "\r\n" +
				"\r\n" +
				string(jsonData)

			syscall.Write(fd, []byte(response))
		} else {
			response := "HTTP/1.1 404 Not Found\r\n" + "\r\n"
			syscall.Write(fd, []byte(response))
		}
	} else {
		response := "HTTP/1.1 404 Not Found\r\n" + "\r\n"
		syscall.Write(fd, []byte(response))
	}
}

func sendAudioRequest(fd int) {
	file, err := os.Open("audio.mp3") // For read access.
	if err!= nil {
		log.Print("error opening file: ", err)
		return
	}

	defer file.Close()

	var requestBody bytes.Buffer
	
	multipartWriter := multipart.NewWriter(&requestBody) // For multipart/form-data
	part, err := multipartWriter.CreateFormFile("audio", "audio.mp3") 
	if err != nil {
		log.Print("error creating form file: ", err)
		return
	}
	
	_, err = io.Copy(part, file)
	if err!= nil {
		log.Print("error copying file: ", errr)
		return
	}

	fieldWriter, err := multiPartWriter.CreateFormField("normal_field")
	if err != nil {
		log.Print("error creating form file: ", err)
		return
	}

	_, err = fieldWriter.Write([]byte("normal_value"))
	if err!= nil {
		log.Print("error writing to form field: ", err)
		return
	}

	multipartWriter.Close() // Close multipart writer

	req, err = http.NewRequest("POST", url, &requestBody)
	if err!= nil {
		log.Print("error creating request: ", err)
		return
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Content-Length", fmt.Sprintf("%d", requestBody.Len())) // For multipart/form-data
	req.Header.Set("Content-Disposition", "form-data; name=\"audio\"; filename=\"audio.mp3\"")
	
	resp, err := http.DefaultClient.Do(req)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция вычисления длины файла в секундах
// На вход получает путь к файлу
func getAudioFileLength(path string) (float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	var totalDuration float64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !isValidHeader(line) {
			return 0, fmt.Errorf("неверный формат заголовка файла: %w", err)
		}

		// Получение длины кадра
		frameDuration := getFrameDurationFromHeader(line)

		totalDuration += frameDuration
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("ошибка сканирования файла: %w", err)
	}

	return totalDuration, nil
}

func isValidHeader(line string) bool {
	return line[0] == 'f'
}

// Анализ строки заголовка и получение длины кадра в секундах
// На вход принимает заголовок WAV-файла
func getFrameDurationFromHeader(line string) float64 {
	fields := strings.Fields(line)
	if len(fields) < 2 || fields[0] != "f" {
		return 0
	}

	frameCount, _ := strconv.Atoi(fields[1])

	// Значение 1152 нужно для подсчета длины аудио в формате WAV (Для 44,1 кГЦ)
	return float64(frameCount) / 1152.0
}

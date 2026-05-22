package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Достаем переменные из env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	url := os.Getenv("TARGET_URL")

	if url == "" {
		log.Fatal("Переменная TARGET_URL пустая")
	}

	// GET запрос
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Ошибка при запросе: %v\n", err)
		os.Exit(1) // Завершаем программу с кодом ошибки
	}

	// Закрываем соединение
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Сервер вернул плохой статус: %s\n", resp.Status)
		os.Exit(1)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("result.html", bodyBytes, 0644)

	if err != nil {
		fmt.Printf("Ошибка при сохранении файла: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Файл result.html создан с правильной кодировкой.")

	/*
		file, err := os.Open("result.html")
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			return
		}
		defer file.Close()
	*/

}

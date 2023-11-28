package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const expectedToken = "4321" // Здесь ваш константный токен

func main() {
	http.HandleFunc("/archive", handleProcess)
	fmt.Println("Server running at port :8088")
	http.ListenAndServe(":8088", nil)
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса (должен быть POST)
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Получение параметров exp_id и token из запроса
	expID := r.FormValue("exp_id")
	token := r.FormValue("token")
	fmt.Println(expID, token)

	// Проверка наличия токена и его соответствия ожидаемому
	if token == "" || token != expectedToken {
		http.Error(w, "Токены не совпадают", http.StatusForbidden)
		fmt.Println("Токены не совпадают")
		return
	}

	// Отправить статус 200 (OK) после проверок
	w.WriteHeader(http.StatusOK)

	// Выполнение операции в фоновом режиме
	go func() {
		// Имитация задержки выполнения в горутине
		delay := 10
		time.Sleep(time.Duration(delay) * time.Second)

		// Генерация случайного результата
		result := fmt.Sprintf("Проверка в архиве прошла успешно для экспедиции с ID: %s", expID)
		if rand.Intn(2) == 0 {
			result = fmt.Sprintf("Проверка в архиве прошла неуспешно для экспедиции с ID: %s", expID)
		}

		// Вывод результата в консоль
		fmt.Printf("Processed with result: %s\n", result)
	}()
}

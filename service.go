package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	expectedToken = "4321"
	updateURL     = "http://127.0.0.1:8000/expedition/update_async/"
)

type ExpeditionResult struct {
	ExpID  string `json:"exp_id"`
	Result string `json:"result"`
	Token  string `json:"token"`
}

func main() {
	http.HandleFunc("/archive", handleProcess)
	fmt.Println("Server running at port :8088")
	http.ListenAndServe(":8088", nil)
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	expID := r.FormValue("exp_id")
	token := r.FormValue("token")
	fmt.Println(expID, token)

	if token == "" || token != expectedToken {
		http.Error(w, "Токены не совпадают", http.StatusForbidden)
		fmt.Println("Токены не совпадают")
		return
	}

	w.WriteHeader(http.StatusOK)

	go func() {
		delay := 10
		time.Sleep(time.Duration(delay) * time.Second)

		result := "Проверка  для экспедиции в архиве прошла успешно"
		if rand.Intn(2) == 0 {
			result = "Проверка для экспедиции в архиве прошла неуспешно"
		}

		// Отправка результата на другой сервер
		expResult := ExpeditionResult{
			ExpID:  expID,
			Result: result,
			Token:  token,
		}
		fmt.Println("json", expResult)
		jsonValue, err := json.Marshal(expResult)
		if err != nil {
			fmt.Println("Ошибка при маршализации JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodPut, updateURL, bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println("Ошибка при создании запроса на обновление:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при отправке запроса на обновление:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Ответ от сервера обновления:", resp.Status)
	}()
}

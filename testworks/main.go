package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	// Симуляция содержимого HTML-файла
	htmlContent := "<h1>{{.Title}}</h1><p>{{.Content}}</p>"

	// Данные для передачи в шаблон
	data := map[string]interface{}{
		"Title":   "Welcome",
		"Content": "Hello World",
	}

	// Парсим шаблон из строки (симулируем загрузку из файла)
	tmpl, err := template.New("html").Parse(htmlContent)
	if err != nil {
		fmt.Println("Ошибка парсинга шаблона:", err)
		return
	}

	// Выполняем шаблон с записью в буфер
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Ошибка выполнения шаблона:", err)
		return
	}
}

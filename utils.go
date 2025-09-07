package main

func GetMessage(page string) string {
	switch page {
	case "home":
		return "Добро пожаловать на наш мини-сервер!"
	case "about":
		return "Это простое приложение на Go из трёх файлов."
	default:
		return "Страница не найдена. Попробуйте другую страницу."
	}
}

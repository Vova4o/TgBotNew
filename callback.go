package main

func CallBackAnswer(data string) string {
	switch data {
	case "1":
		return "Вы выбрали кнопку 1"
	case "2":
		return "Вы выбрали кнопку 2"
	case "3":
		return "Вы выбрали кнопку 3"
	}
	return ""
}

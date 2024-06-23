## Бот для Telegram на Go

### Описание:

Данный проект представляет собой пример создания бота для Telegram на языке Go.
Он позволяет взаимодействовать с пользователем,
отвечать на сообщения, отправлять кнопки и клавиатуры.

### Возможности:

- **Приветствие пользователя:** При первом запуске бота пользователь получает приветственное сообщение и предложение использовать кнопки для навигации.
- **Главное меню:** Представлено в виде клавиатуры с кнопками, каждая из которых может вести к новому разделу или функционалу бота.
- **Обработка сообщений:** Бот отвечает на сообщения пользователя в соответствии с заданной логикой.
- **Кнопки с ответом:** Бот может отправлять сообщения с кнопками, нажатие на которые приводит к новым действиям.
- **Сохранение информации о пользователях:** Бот может сохранять информацию о пользователях в базе данных MySQL.

### Структура проекта:

- **main.go:** Содержит код основной программы, включая функции подключения к Telegram и MySQL, обработки обновлений, взаимодействия с пользователем.
- **dbConnect.go:** Содержит функции подключения к MySQL.
- **tgConnect.go:** Содержит функции подключения к Telegram.
- **keyboard.go:** Содержит функции создания клавиатур для бота.
- **markup.go:** Содержит константы с клавиатурами для бота.
- **utils.go:** Содержит вспомогательные функции.

### Первый запуск:

1. **Скачайте проект:** Скачайте архив с кодом проекта и распакуйте его в удобную для вас директорию.
2. **Создайте файл .env:** Создайте файл ```.env``` в корневой папке проекта и добавьте в него следующие строки:
```
BOT_TOKEN=<your_bot_token>
DB_USER=<mysql_username>
DB_PASSWORD=<mysql_password>
DB_NAME=<mysql_database_name>
```
3. **Установите зависимости:** В терминале перейдите в директорию проекта и выполните команду ```go get.```
4. **Запустите бота:** В терминале выполните команду ```go run main.go.```

### Использование:

- **Запустите бота.**
- **Нажмите кнопку "start"** в главном меню.
- **Взаимодействуйте с ботом**, используя кнопки и текстовые сообщения.

### Важные моменты:

- **Для работы бота требуется подключение к Telegram.** Получите бот-токен в Telegram BotFather и укажите его в файле ```.env.```
- **Для хранения информации о пользователях требуется MySQL.** Укажите данные подключения к MySQL в файле ```.env.```
- **Данный проект является базовым примером.** Вы можете его доработать, добавив новые функции, клавиатуры, логику обработки сообщений.
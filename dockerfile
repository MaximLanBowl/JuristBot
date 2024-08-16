# Используем официальный образ Golang в качестве базового
FROM golang:1.22.3

# Создаем рабочую директорию внутри контейнера
WORKDIR /app
# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем модули Go
RUN go mod download

# Копируем остаток исходного кода в контейнер
COPY . .

# Компилируем Go-приложение
RUN go build -o main ./cmd/main.go

# Указываем команду запуска контейнера
CMD ["./main"]
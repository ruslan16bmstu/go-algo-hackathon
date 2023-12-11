FROM golang:1.21 AS builder

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Сборка проекта
RUN CGO_ENABLED=0 go build -o app

# Финальный этап
FROM debian:bullseye-slim

WORKDIR /app

# Копируем исполняемый файл из этапа сборки
COPY --from=builder /app/app .

COPY main-server.config.toml .
ENV CONFIG=/app/main-server.config.toml
COPY ./db/stocks.json ./db/
COPY ./db/predictions.json ./db/

# Порт, который будет использоваться приложением
EXPOSE 8000

# Команда для запуска приложения внутри контейнера
CMD ["./app"]

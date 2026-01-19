# LibraryManager

> RESTful API для управления личной библиотекой на Go

## О проекте

LibraryManager — это HTTP-сервис для учёта прочитанных книг, написанный на чистом Go с использованием best practices. Проект демонстрирует навыки проектирования REST API, работы с конкурентностью и организации чистой архитектуры.

## Что реализовано

### Архитектура
- **Чистое разделение слоёв**: бизнес-логика (`library`), HTTP-слой (`http`), DTO
- **Thread-safe операции**: использование `sync.RWMutex` для безопасной работы с данными
- **Валидация на уровне domain**: проверка корректности данных при создании книг
- **Правильная обработка ошибок**: типизированные ошибки с корректными HTTP-статусами

### API endpoints

```
POST   /books              - Добавить книгу
GET    /books              - Получить список книг (с фильтрацией)
GET    /books/{title}      - Получить книгу по названию
PATCH  /books/{title}/finish - Отметить книгу как прочитанную
DELETE /books/{title}      - Удалить книгу
```

### Фичи

**CRUD операции** для книг  
**Фильтрация** по автору и статусу прочтения  
**Отслеживание времени** добавления и завершения чтения  
**Concurrent-safe** работа с данными  
**REST-compliant** HTTP статусы и методы  
**JSON** serialization/deserialization  

## Технологический стек

- **Go** — язык программирования
- **gorilla/mux** — популярный HTTP роутер с поддержкой path variables
- **sync.RWMutex** — для thread-safe доступа к данным
- **encoding/json** — работа с JSON

## Запуск

```bash
go mod download
go run main.go
```

Сервер запустится на `http://localhost:9091`

## Примеры использования

### Добавить книгу
```bash
curl -X POST http://localhost:9091/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Чистый код",
    "author": "Роберт Мартин",
    "numberOfPages": 464
  }'
```

### Получить все книги автора
```bash
curl "http://localhost:9091/books?author=Роберт%20Мартин"
```

### Получить только прочитанные книги
```bash
curl "http://localhost:9091/books?isFinished=true"
```

### Отметить книгу прочитанной
```bash
curl -X PATCH http://localhost:9091/books/Чистый%20код/finish
```

### Удалить книгу
```bash
curl -X DELETE http://localhost:9091/books/Чистый%20код
```

## Структура проекта

```
LibraryManager/
├── main.go              # Точка входа
├── library/             # Domain layer
│   ├── library.go       # Основная бизнес-логика
│   ├── book.go          # Модель и конструктор
│   └── errors.go        # Типизированные ошибки
└── http/                # HTTP layer
    ├── server.go        # Настройка сервера и роутинга
    ├── handlers.go      # HTTP handlers
    ├── helpers.go       # Вспомогательные функции
    ├── middlewares.go   # HTTP middlewares
    └── dto/             # Data Transfer Objects
        ├── book.go      # DTO для книги
        └── err.go       # DTO для ошибок
```

## Технические решения

### Concurrency Safety
Используется `sync.RWMutex` для безопасной работы нескольких горутин с хранилищем книг. Read-операции используют `RLock()` для параллельного чтения, write-операции блокируют полностью.

### Обработка ошибок
Все доменные ошибки типизированы и обрабатываются с правильными HTTP-статусами:
- `404 Not Found` — книга не найдена
- `409 Conflict` — книга уже существует / уже прочитана
- `400 Bad Request` — невалидные данные

### Валидация
Валидация происходит на уровне создания domain-объектов, что гарантирует невозможность создания некорректных сущностей.
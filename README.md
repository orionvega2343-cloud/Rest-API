# Blog API

REST API для блога с постами и комментариями. Написан на Go с использованием PostgreSQL.

## Технологии

- Go
- PostgreSQL
- `net/http`
- `database/sql` + `github.com/lib/pq`

## Запуск

### 1. Требования

- Go 1.21+
- PostgreSQL

### 2. Создай базу данных

```sql
CREATE DATABASE blog;
```

### 3. Настрой переменную окружения

```bash
# Windows PowerShell
$env:DB_PASSWORD="твойпароль"

# Linux / macOS
export DB_PASSWORD=твойпароль
```

### 4. Запусти сервер

```bash
go run .
```

Сервер запустится на `http://localhost:8080`

---

## Эндпоинты

### Посты

| Метод | URL | Описание |
|---|---|---|
| GET | `/posts` | Получить все посты |
| POST | `/posts` | Создать пост |
| PUT | `/posts?id=1` | Обновить пост |
| DELETE | `/posts?id=1` | Удалить пост |
| GET | `/post?id=1` | Получить пост по ID |

### Комментарии

| Метод | URL | Описание |
|---|---|---|
| GET | `/comments?id=1` | Получить комментарии поста |
| POST | `/comments?id=1` | Добавить комментарий к посту |

---

## Примеры запросов

### Получить все посты

```bash
curl http://localhost:8080/posts
```

### Создать пост

```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"Заголовок","text":"Текст поста","author":"Иван","date":"18.04.2026"}'
```

### Обновить пост

```bash
curl -X PUT "http://localhost:8080/posts?id=1" \
  -H "Content-Type: application/json" \
  -d '{"title":"Новый заголовок","text":"Новый текст","author":"Иван","date":"18.04.2026"}'
```

### Удалить пост

```bash
curl -X DELETE "http://localhost:8080/posts?id=1"
```

### Получить комментарии поста

```bash
curl "http://localhost:8080/comments?id=1"
```

### Добавить комментарий

```bash
curl -X POST "http://localhost:8080/comments?id=1" \
  -H "Content-Type: application/json" \
  -d '{"text":"Отличный пост!","author":"Мария"}'
```

---

## Структуры данных

### Post

```json
{
  "id": 1,
  "title": "Заголовок",
  "text": "Текст поста",
  "author": "Иван",
  "date": "18.04.2026"
}
```

### Comment

```json
{
  "id": 1,
  "text": "Отличный пост!",
  "author": "Мария",
  "post_id": 1
}
```

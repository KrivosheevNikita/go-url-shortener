# 🌐 Go URL Shortener

A simple URL shortening service built with Go, PostgreSQL, Redis, and Docker.



## 🛠️ Tech Stack

- Go (Golang)
- PostgreSQL
- Redis
- Gorilla Mux
- Docker

## ✨ Features

- 🔗 Create short URLs with custom or auto-generated IDs
- 📊 Track usage statistics for link
- 🕒 Set expiration date for link
- ⚡ Fast redirection using Redis cache
- 🐳 Dockerized for easy deployment



## ⚙️ Setup

### 1. Clone the repository

```bash
git clone https://github.com/KrivosheevNikita/go-url-shortener.git
cd go-url-shortener
```
### 2. Set up environment variables
Create a .env file in the root directory with the following content:
```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=urlShortener
POSTGRES_PORT=5432

POSTGRES_ADDR=postgres://postgres:password@postgres-shortener:5432/urlShortener?sslmode=disable

REDIS_ADDR=redis-shortener:6379
```
### 3. Run the App

Start using Docker Compose:

```bash
docker-compose up --build
```
## 📦 API Endpoints

### POST /shorten

Create a new short URL.

#### Request Body
```json
{
  "url": "https://example.com",
  "expiry": "2025-12-31T23:59:59Z", 
  "custom_id": "123"
}
```
expiry - optional

custom_id - optional

#### Response

- `200 OK`
```json
{
  "short_url": "http://localhost:8080/123"
}
```

#### Errors

- `400 Bad Request` – Invalid input.
- `409 Conflict` – ID already exists.
- `500 Internal Server Error` – Server error.

---

### GET /{id}

Redirect to the original URL.

#### Response

- `302 Found` – Redirect to the original URL.


#### Errors

- `404 Not Found` – The short URL doesn't exist or has expired.
- `500 Internal Server Error` – Server error.

---

### DELETE /delete/{id}

Delete a short URL.

#### Response

- `204 No Content` – URL deleted successfully.

#### Errors

- `404 Not Found` – The short URL does not exist.
- `500 Internal Server Error` – Server error.

---

### GET /stats/{id}

Get statistics for a short URL (usage count, expiry, etc.).

#### Response

- `200 OK`
```json
{
  "url": "https://example.com",
  "expiry": "2025-12-31T23:59:59Z",
  "usage_count": 12
}
```

#### Errors

- `404 Not Found` – The short URL does not exist or has expired.
- `500 Internal Server Error` – Server error.


# 🚀 URL Shortener Service (Golang)

A production-ready URL shortener service built with Go, featuring custom aliases, expiration, analytics, rate limiting, and Dockerized environment.

---

## 📌 Overview

This project is a backend service that transforms long URLs into short, shareable links. It is designed using clean architecture principles and simulates real-world backend systems with database integration, analytics tracking, and containerization.

---

## ✨ Features

### 🔗 Core Features

* Shorten long URLs
* Custom alias (e.g. `/my-link`)
* Redirect to original URL
* Expiration support

### 📊 Analytics

* Track total clicks
* Store:

  * IP Address
  * User-Agent
  * Timestamp

### 🚦 Rate Limiting

* IP-based rate limiting middleware
* Prevent abuse/spam

### ⚙️ Migration System

* Auto-run database migration on app start
* Versioned SQL files

### 🐳 Docker Support

* Fully containerized app
* Connects to existing PostgreSQL container
* Easy setup with Docker Compose

---

## 🧱 Tech Stack

* Golang (Gin)
* PostgreSQL
* Docker & Docker Compose
* SQL (raw queries)
* Clean Architecture

---

## 📁 Project Structure

```
url-shortener/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── middleware/
│
├── migrations/
│   ├── 001_init.up.sql
│   └── 001_init.down.sql
│
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## ⚙️ Setup & Installation

### 1. Clone Repository

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

---

## 🐳 Running with Docker (Recommended)

### 🔹 Prerequisites

* Docker installed
* PostgreSQL container already running

---

### 🔹 Run Application

```bash
docker compose up --build
```

---

### 🔹 Access

* API → http://localhost:8080
* pgAdmin → http://localhost:8082 (if running)

---

## 🧪 Running Locally

### 1. Setup PostgreSQL

```sql
CREATE DATABASE url_shortener;
```

---

### 2. Update DB Config

Edit file:

```
internal/repository/db.go
```

Change host:

```go
host=localhost
```

---

### 3. Run App

```bash
go run cmd/main.go
```

---

## 🔌 API Endpoints

### ➕ Create Short URL

```http
POST /shorten
```

**Request Body:**

```json
{
  "url": "https://example.com",
  "custom_alias": "my-link",
  "expired_at": "2026-12-31T23:59:59Z"
}
```

---

### 🔁 Redirect

```http
GET /:code
```

---

### 📊 Get Statistics

```http
GET /stats/:code
```

---

## 🗄️ Database Schema

### urls

* id
* original_url
* short_code
* clicks
* created_at
* expired_at

### url_clicks

* id
* short_code
* clicked_at
* user_agent
* ip_address

---

## 🔥 Engineering Highlights

* Clean Architecture (Handler → Service → Repository)
* Docker-based system
* Auto database migration
* Rate limiting middleware
* Async analytics logging
* Container networking (service-to-service)

---

## 🚀 Future Improvements

* Redis caching (low latency redirect)
* Distributed rate limiting
* API authentication (API Key / JWT)
* Dashboard UI
* Link preview (OpenGraph)

---

## 📄 License

MIT License

---

## 👨‍💻 Author

Alfarizi 🚀

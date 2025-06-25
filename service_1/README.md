
### 📘 Service 1 – Golang API

This is a lightweight HTTP service written in Go. It serves two endpoints: `/ping` and `/hello`. It is designed to be used behind a reverse proxy like Nginx.

---

### 🔧 Endpoints

| Method | Path     | Description      | Response Example                      |
| ------ | -------- | ---------------- | ------------------------------------- |
| GET    | `/ping`  | Health check     | `{"status": "ok", "service": "1"}`    |
| GET    | `/hello` | Greeting message | `{"message": "Hello from Service 1"}` |

---

### 🚀 Running Locally

#### Prerequisites

* Go 1.22+
* Docker (optional, for containerized setup)

#### Run with Go:

```bash
go run main.go
```

This will start the server on port `8001`.

---


Then visit:

* [http://localhost:8001/ping](http://localhost:8001/ping)
* [http://localhost:8001/hello](http://localhost:8001/hello)

---

### 📂 Project Structure

```
service1/
├── main.go
└── README.md
```

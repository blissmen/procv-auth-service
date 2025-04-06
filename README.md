# 🔐 Auth Service – ProCV Authentication Microservice (Go)

The **Auth Service** is a lightweight Go-based microservice within the ProCV platform. It handles user registration, login, and JWT-based authentication for secure communication across microservices.

This service was scaffolded to remain minimal, with a focus on configurability and future extensibility.

---

## 📦 Structure

```bash
auth-service/
├── config/
│   └── config.yaml         # Service configuration (port, secrets)
├── main.go                 # Entry point
├── go.mod                  # Module file
└── README.md
```

---

## 🎯 Responsibilities

- 🔐 Register new users
- 🔑 Handle secure login with password hashing
- 🪪 Generate and validate JWT tokens
- 📬 Interface with Gateway to authenticate requests

---

## ⚙️ Configuration

Service settings are stored in `config/config.yaml`

### Example:
```yaml
server:
  host: 0.0.0.0
  port: 8001

jwt:
  secret: "your_super_secret_key"
  expiration_minutes: 60
```

---

## 🚀 Usage

### 1. Clone the repo
```bash
git clone https://github.com/blissmen/procv-auth-service.git
cd procv-auth-service
```

### 2. Build and Run
```bash
go run main.go
```

---

## 🧪 Endpoints (planned)

- `POST /register` – Register a new user
- `POST /login` – Authenticate user and return token
- `GET /verify` – Validate token for protected routes

---

## 🔐 Tech Stack

- **Go**
- **Gin or net/http** – HTTP routing (based on your implementation)


---

## 📜 License
[MIT License](LICENSE)

---

## 👨‍💻 Author
Built by **Tangwe Caleb Ojang**

# 🏗️ Ports and Adapters Architecture

This project is a **Go backend scaffold** following **Clean Architecture / Hexagonal Architecture** principles.  
It demonstrates how to structure a project with clear separation of **Domain**, **Application**, and **Infrastructure**, while using:

- [Echo](https://echo.labstack.com/) for HTTP server.  
- [Uber Dig](https://pkg.go.dev/go.uber.org/dig) for Dependency Injection.  

---

## 📂 Project Structure

```bash
├── cmd/
│ └── api/ # Entrypoint for HTTP server
│   ├── di/ # Dig container (composition root)
│   ├── httpserver/ # Middlewares, routes and Server definition
│   └── main.go # Init Dig container and start Echo server, main function
│
├── docs/ # README, diagrams and any documentation
│
├── internal/
│ ├── Infrastructure/ # Adapters (implementations)
│ │ ├── http/ # Echo HTTP handlers
│ │ │ └── user_handler.go
│ │ ├── data/ # Repositories (e.g., SQL)
│ │ └── messaging/ # (Optional) Kafka, RabbitMQ, etc.
│ │
│ ├── application/ # Use cases (business logic orchestration)
│ │ └── user_service.go
│ │
│ ├── config/ # Config structs and env loader 
│ │
│ └── domain/ # Models, value objects, business validations, ports (interfaces)
│   ├── interfaces/
│   ├── models/
│   └── validations/
│
├── pkg/ # Reusable modules (e.g. Logger)
│
└── go.mod
```

---

## 🧩 Layers

- **Domain** → Enterprise business rules  
  - Models, value objects, interfaces (ports), pure logic.  
  - No external dependencies.  

- **Application** → Application business rules  
  - Orchestrates use cases (e.g. `UserService.Create`).  
  - Depends only on **domain interfaces**.  

- **Infrastructure** → Implementation details(adapters)  
  - HTTP handlers, Data repositories, messaging.  
  - Implement the interfaces defined in **domain**.  

- **Entrypoint (`cmd/`)** → Composition root  
  - Wires dependencies using `dig`.  
  - Starts Echo HTTP server.  

---

## 🚀 Getting Started

### 1. Clone and install

```bash
git clone git@github.com:IsaiasGC/go-ports-adapters-scaffold.git
cd go-ports-adapters-scaffold
go mod tidy
```

### 2. Run the API

```bash
go run ./cmd/api
```

Server starts on `http://localhost:8080`.

---

## 📡 Example Request

### Create User

```bash
curl http://localhost:8080/health
```

### Response

```json
{
    "status": "pass",
    "version": "0.1.0",
    "checks": [
        {
            "componentName": "ports-adapters",
            "componentType": "pod",
            "time": "0.000ms",
            "status": "pass"
        }
    ]
}
```

## 🔑 Key Concepts

- DI with Dig → Only the cmd/ layer knows about the DI container. Other layers just expose constructors (NewXxx).

- Testability → You can unit test domain and application without HTTP or DB.

- Flexibility → You can add another entrypoint (cmd/worker, cmd/migrator) with different wiring but same application/domain.

## 🛠️ Next Steps

- Add middlewares (logging, auth) in adapter/http.

---

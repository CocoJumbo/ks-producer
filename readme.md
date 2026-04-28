# Kafka Sandbox Producer (Learning Project)

This project is a **simple Kafka producer written in Go**, built for learning purposes.
It demonstrates how to send messages to Apache Kafka using a clean, modular structure.

It is part of a bigger setup:
👉 https://github.com/CocoJumbo/ks-consumer

---

## 📌 Purpose

This project is designed to help you:

* Learn Go project structure (`cmd/`, `internal/`)
* Understand Kafka basics
* Build a simple HTTP API that sends messages to Kafka

---

## 🧱 Project Structure

```
cmd/
  main.go              # application entry point

internal/
  api/
    handler.go         # HTTP endpoints (Gin)
  kafka/
    producer.go        # Kafka producer implementation
  config/
    config.go          # configuration (env-based)
```

---

## 🚀 How It Works

Flow:

```
HTTP request → Gin API → Kafka Producer → Kafka Topic
```

Example:

```
POST /simple-messages
{
  "data": "hello kafka"
}
```

👉 Message is sent to Kafka topic:

```
simple-topic
```

---

## ⚙️ Requirements

* Go 1.21+
* A running Kafka instance (reuse from the consumer project)

---

## ▶️ Run the Producer

### 1. Install dependencies

```bash
go mod tidy
```

---

### 2. Generate Swagger docs (IMPORTANT)

Before generating Swagger files, remove old ones to avoid version conflicts:

```bash
rm -rf docs
```

Windows (PowerShell):

```powershell
Remove-Item -Recurse -Force docs
```

Then generate fresh docs:

```bash
swag init --generalInfo cmd/main.go --output docs
```

> Run this step whenever:
>
> * you change API annotations
> * you update Swagger dependencies


### 3. Run application

```bash
go run cmd/main.go
```

Server starts on:

```
http://localhost:8081
```

---

## 📬 API Usage

### Endpoint

```
POST /simple-messages
```

### Request body

```json
{
  "data": "hello kafka"
}
```

### Response

```json
{
  "status": "sent"
}
```

---

## 📊 Swagger UI

After starting the app:

```
http://localhost:8081/swagger/index.html
```

---

## 🧠 Kafka Notes

* Topic used: `simple-topic`
* No message key → messages are distributed randomly across partitions
* Uses pure Go client (no native dependencies)

---

## 🔗 Related Project

Consumer (Spring Boot):

👉 https://github.com/CocoJumbo/ks-consumer

---

## 💡 Next Steps

You can extend this project with:

* Kafka consumer in Go
* message keys (partition control)
* retries & error handling
* dead-letter topics (DLQ)
* authentication (SASL/SSL)

---

## 🧪 Learning Focus

This project is intentionally simple and focuses on:

* understanding Kafka concepts
* learning Go idioms
* experimenting safely

---

## 📄 License

MIT (or your preferred license)

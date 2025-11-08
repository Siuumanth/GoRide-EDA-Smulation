
# **GoRide: In-Memory Auto-Scaling Event-Driven Model**

GoRide is a high-throughput simulation built to explore and master Go's concurrency primitives by modeling a simplified ride-sharing dispatch and fulfillment system.

It implements a fully **in-memory Event-Driven Architecture (EDA)** where services communicate asynchronously through Go channels — showcasing goroutines, channels, contexts, mutexes, and dynamic scaling in action.

---

## 🚀 **Architecture & Core Principles**

![](https://github.com/Siuumanth/GoRide-EDA-Smulation/blob/main/images/1.jpg?raw=true)

GoRide operates as a **true concurrent system**, composed of **decoupled service workers** coordinated by a dynamic control plane.

### **Key Components**

* **🧠 EventBus (`events/`)**
  The central event router responsible for distributing events (e.g., `TripRequestedEvent`, `PaymentEvent`) from publishers to all subscribed services.

* **⚙️ Service Workers (`services/`)**
  Independent, concurrent components — **Driver**, **Trip**, **Payment**, **PaymentAsk**, **Notification** and **TripCompleted** — each running as separate goroutines handling domain-specific logic.

* **📈 AutoScaler (`core/AutoScaler`)**
  A custom-built load manager that dynamically scales worker pools up or down based on EventBus queue size and system activity.

---

## Project Structure:

```bash
goride/
├── go.mod
├── go.sum
│
├── main.go               # Entry point: Initializes EventBus, AutoScaler, and starts simulation
│
├── core/
│   ├── AutoScaler.go     # Contains the AutoScaler logic (scaling up/down)
│   ├── Pubs&Subs.go       # Defines the publishers and subscribers logic
│   └── workerpool.go     # Manages the service worker pools and handles events
│
├── events/
│   ├── events.go         # Defines all Event structs (e.g., TripCreated, DriverMatched)
│   ├── eventBus.go       # Event dispatching logic
|
├── services/             # Defines independant services
│   ├── trip_service.go
│   ├── driver_service.go
│   ├── payment_service.go
│   └── ... (etc.)
│
├── users.go # Logic to simulate users and publish events
│
└── notifications.log     # Example output log file
```

---

## 🧵 **Concurrency Pillars**

* **Goroutines & Channels** – Enable asynchronous task execution and inter-service communication.
* **Contexts (`context.Context`)** – Provide graceful cancellation and cascading shutdown of services.
* **Mutex (`sync.Mutex`)** – Ensure thread safety when accessing shared global resources (e.g., available driver lists).
* **Graceful Shutdown** – Implements idle-based exit logic that automatically terminates the application when the system remains inactive for a set duration.

---

## ⚙️ **Key Technical Features**

* **Dynamic Worker Pools**
  Automatically adjusts concurrency level based on the **EventBus backlog**, scaling up under load and scaling down when traffic decreases.

* **Bottleneck Resolution**
  Simulates performance tuning to fix bottlenecks caused by chained events or unbalanced service delays.

* **In-Memory Pub/Sub**
  Implements a lightweight asynchronous communication system using Go channels, mimicking Kafka-like event-driven messaging.

* **Race Condition Prevention**
  Uses `sync.Mutex` to safely manage shared state, ensuring deterministic and error-free execution under high concurrency.
  
---

## 🧩 **Simulation Output**

When executed, the system logs all simulated events to a `notifications.log` file in the project root.
You’ll see logs for trip requests, driver matches, payments, and shutdown sequences.

![](https://github.com/Siuumanth/GoRide-EDA-Smulation/blob/main/images/2.png?raw=true)

---

### Console Output:
![](https://github.com/Siuumanth/GoRide-EDA-Smulation/blob/main/images/3.png?raw=true)

---

## ▶️ **Getting Started**

### **Prerequisites**

* Go **1.18+** installed

### **Run the Simulation**

```bash
go run main.go 
```

> 💡 *Depending on your Go version, you may simply run:*

```bash
go run .
```

---

## 🧠 **Summary**

GoRide demonstrates how Go’s concurrency model can power a realistic, event-driven system with dynamic scaling — all in memory, without external dependencies.
It’s a practical exploration of **concurrent system design**, **load management**, and **graceful orchestration** using pure Go.

---



# High-Scale URL Shortener (Go + Redis)

A high-performance, low-latency URL shortening service built with **Golang**, designed for scalability and efficiency. This service leverages **Redis** for $O(1)$ lookup times and implements a **Bloom Filter** to optimize database throughput.



## Key Features

* **Sub-10ms Redirection:** Leverages Redis in-memory storage for near-instantaneous URL resolution.
* **Bloom Filter Optimization:** Implemented a probabilistic data structure (Bloom Filter) to perform membership testing, reducing unnecessary Redis lookups by **40%** and preventing cache penetration.
* **Deterministic Hashing:** Uses **SHA-256** hashing combined with **Base62 encoding** to generate unique, URL-friendly short aliases.
* **Clean Architecture:** Modular code structure separating handlers, storage logic, and utility functions for high maintainability.

---

## Tech Stack

* **Language:** Go (Golang)
* **Web Framework:** Gin Gonic
* **Database:** Redis (In-memory Key-Value store)
* **Data Structures:** Bloom Filter (bits-and-blooms)
* **API Protocol:** REST

---

## System Architecture

1.  **Request Layer:** Gin handles incoming POST (creation) and GET (redirection) requests.
2.  **Validation Layer:** A **Bloom Filter** checks if a short URL potentially exists before hitting the database.
3.  **Storage Layer:** Redis stores the mapping between the short alias and the original URL with a configurable TTL.



---

## Getting Started

### Prerequisites
* Go 1.2x+
* Redis Server (Running on `localhost:6379`)

### Installation
1. Clone the repository:
   ```bash
   git clone [https://github.com/nandani2203/url_shortener.git](https://github.com/nandani2203/url_shortener.git)
   cd url_shortener

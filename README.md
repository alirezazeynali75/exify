# Exify Payment Microservice

**Exify** is a microservice written in Go designed to handle payment transactions (deposits and withdrawals) using multiple payment gateways. The service is designed for extensibility, allowing for easy integration of new gateways. Currently, it integrates two types of payment gateways - JSON over HTTP and SOAP/XML over HTTP. It uses **MySQL** as the database and **Redpanda** (Kafka-compatible streaming platform) for asynchronous messaging.

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Design Decisions](#design-decisions)
- [Future Improvements](#future-improvements)

---

## Overview

This microservice provides:
- **Deposit (Cash-In)** and **Withdrawal (Cash-Out)** handling.
- Integration with two payment gateways:
  - **Gateway A**: Uses JSON over HTTP.
  - **Gateway B**: Uses SOAP/XML over HTTP.
- **Asynchronous Callbacks** to handle transaction status updates.
- **Extensibility** to support new payment gateways in the future.
- **Resilience and Fault Tolerance** using circuit breakers, retries, and timeout mechanisms.
- **Logging and Error Handling** to ensure secure and clear transaction processing.

---

## Architecture

The **Exify** microservice is built using:
- **Go** as the main programming language.
- **MySQL** as the primary database for storing transaction details.
- **Redpanda** for handling event-driven asynchronous messaging.
- A **modular architecture** to ensure extensibility and maintainability.

### Components
- **Core Service**: Contains business logic for handling transactions.
- **Payment Gateways**: Gateway A and Gateway B integrations, designed to be easily extended to support new gateways.
- **Repositories**: Handles data persistence in MySQL.
- **Asynchronous Processing**: Uses Redpanda for handling callbacks and event-based messaging.
- **Configuration Management**: Environment-based configurations for different environments (e.g., development, production).


---

## Setup Instructions

### Prerequisites
- Go 1.22+
- Docker (for containerization and running MySQL/Redpanda)
- MySQL Database
- Redpanda (Kafka-compatible platform)

### Building and Running the Service

**Clone the repository:**
```bash
git clone https://github.com/yourusername/exify.git
cd exify
```

Using the Makefile
Install Dependencies:

```bash
make install-dependencies
```

**Build the Service**:

```bash
make build && make build-cli
```

**Run the Service**:

```bash
./exify-cli migrate up && ./exify
```

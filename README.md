**GoCommerce** is a simple, production-ready e-commerce backend template built with **Go** and the **Echo** web framework. It provides RESTful APIs for managing products, carts, orders, customers, and payments. Designed as a starter kit for building scalable e-commerce services or for learning how to structure Go web applications.

## Features
- RESTful API design with Echo
- Product CRUD (create, read, update, delete)
- Cart and checkout flow
- Order management and status tracking
- User authentication (JWT)
- Payment integration (Stripe example)
- PostgreSQL (GORM) for persistence
- Environment based configuration
- Docker & docker-compose for local development
- Unit & integration test scaffolding

## Tech Stack
- Language: Go (1.XX+)
- Web framework: Echo
- Database: PostgreSQL (GORM)
- Payments: Stripe (optional)
- Auth: JWT
- Migrations: golang-migrate (or embed SQL)
- Dependency management: Go modules
- Containerization: Docker

## Quickstart (Development)
1. Clone the repo:
   ```bash
   git clone https://github.com/minulhasanrokan/go-ecommerce.git
   go-ecommerce

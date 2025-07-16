##  Tech Stack

This project follows a client-server architecture with separate layers for frontend, backend API, and an AI service.

| Layer         | Technologies                                 | Purpose                                                    |
|---------------|----------------------------------------------|------------------------------------------------------------|
| Frontend      | Flutter 3.32.4 (Dart 3.8.1)                  | Cross-platform UI for test creation and result display     |
| Backend (API) | Go 1.24, Clean Architecture                  | REST API, routing, business logic                          |
| AI Service    | Python 3.11, FastAPI, OpenAI SDK             | LLM-based test generation and natural language processing  |
| DevOps        | Docker, Docker Compose, GitHub Actions       | Containerization, local development, CI/CD                 |
| Hosting       | VPS (Ubuntu 22.04), nginx                    | Production deployment and reverse proxy                    |
| Testing       | Go test, Pytest, Flutter test                | Unit, integration, and end-to-end testing                  |

---

###  Technology Rationale

#### **Go (Backend API)**
- High performance and static binaries simplify deployment.
- Clean Architecture enforces separation of concerns and testability.
- Built-in support for concurrency and scalability.

#### **Python + FastAPI (AI Layer)**
- Easy integration with OpenAI SDK and external APIs.
- FastAPI offers high performance with simple async support.
- Ideal for LLM-based services and prototyping.

#### **Flutter (Frontend)**
- Single codebase for web and mobile.
- Rapid UI development with strong test support.
- Dart offers good performance and modern tooling.


#### **Docker + GitHub Actions**
- Consistent environments across dev and prod.
- Automated CI/CD pipelines, testing, and deployment.


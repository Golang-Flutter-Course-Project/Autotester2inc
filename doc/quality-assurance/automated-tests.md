
#  Automated Tests

This document outlines the automated testing strategy and test types used in the project. Automated tests help ensure code quality, prevent regressions, and speed up development.

---

##  Types of Tests

| Layer        | Language | Type             | Description                                 |
|--------------|----------|------------------|---------------------------------------------|
| Backend (Go) | Go       | Unit & Integration | Covers core logic, HTTP handlers          |
| Frontend     | Dart     | Widget & Unit    | UI interaction tests and business logic     |
| AI Service   | Python   | Unit & HTTP      | Validates test runner and OpenAI responses  |
---

##  Backend Tests (Go)

- Uses the standard `testing` package
- Located in `*_test.go` files
- Covers:
  - HTTP handlers (`/api/results`)
  - Validation logic
  - Cookies
  - Middleware

**Run all backend tests:**
```bash
cd backend
go test ./...
```

## Frontend Tests (Flutter)

- Unit test for widgets loading implemented
- Located in `test` folder
- Cover:
  - Home page
  - Result page
  - Loading page

**Run all frontend tests:**
```bash
cd frontend
flutter test
```

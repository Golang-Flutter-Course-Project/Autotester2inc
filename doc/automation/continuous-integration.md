# Continuous Integration Guide

This document describes the automated CI pipeline for the project, including the tools used and steps executed on each commit and pull request.

---

##  Overview

Our CI system ensures that every code change is automatically built, tested, and validated before being merged into the main branch.

CI is powered by **GitHub Actions** and defined in the following files:

- `.github/workflows/flutter_ci.yml`
- `.github/workflows/go-ci.yml`
- `.github/workflows/python-ci.yml`

---

## Tools Used

| Tool            | Purpose                          |
|-----------------|----------------------------------|
| GitHub Actions  | CI runner                        |
| Docker          | Containerized build/test         |
| go test,golangCI| Backend unit tests and linting   |
| Flutter         | Frontend build & widget testing  |
| pytest          | Python backend testing           |

---

##  Pipeline Steps

###  Backend (Go)
1. Run `go mod tidy`
2. Run `go test ./...`

###  Frontend (Flutter)
1. Run `flutter pub get`
2. Run `flutter test`

###  Python (LLM Service)
1. Run `black --check`
2. Run `isort --check-only`
3. Run `pytest`

---

### Continuous integration

Flutter:
- Used in-build flutter tools for testing
- Implemented unit-tests for widgets
- `frontend/test` contains all tests

Flutter CI:
- Link to CI: https://github.com/cQu1x/Autotester/blob/main/.github/workflows/flutter_ci.yml
- Downloads flutter
- Runs tests
- Creates build for web application (important for code updates)

Golang:
 - Used in-build golang tools for testing
 - Implemented unit test for cookies and validators. Implemented integration tests for handlers
 - `tests/` contains all tests

Golang CI:
 - Link to CI: https://github.com/cQu1x/Autotester/actions/workflows/go-ci.yml
 - Downloads Golang
 - Runs tests and linting

Python:
 - Used python libraries (pytest) for testing
 - Implemented tests for parser and LLM-connection
 - `/test_1.py` contains all tests

Python CI:
 - Link to CI: https://github.com/cQu1x/Autotester/actions/workflows/python-ci.yml
 - Downloads Python
 - Runs tests

##  Triggers

CI workflows run on:

- Every pull request to `main, go-backend, fontend_branch`
- Manual dispatch via GitHub UI

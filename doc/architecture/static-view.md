graph TD
    User["Пользователь"]
    Frontend["Frontend (Flutter Web/App)"]
    Backend["Backend (Go API)"]
    DB["PostgreSQL Database"]
    ExtService["Внешние сервисы (например, сайты для проверки URL)"]

    User -- "HTTP/HTTPS" --> Frontend
    Frontend -- "REST API" --> Backend
    Backend -- "SQL" --> DB
    Backend -- "HTTP/HTTPS" --> ExtService

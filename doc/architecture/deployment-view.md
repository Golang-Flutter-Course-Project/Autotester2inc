## Deployment
``` mermaid
graph TD
    subgraph Frontend
        FE["Flutter App"]
    end

    subgraph Backend
        API["Go REST API"]
        Auth["Auth Module"]
        DBLayer["Database Layer"]
        Routes["Routes/Handlers"]
        Middleware["Middleware"]
    end

    subgraph Database
        PG["PostgreSQL"]
    end

    FE -- "REST API" --> API
    API -- "Auth Calls" --> Auth
    API -- "DB Queries" --> DBLayer
    API -- "Route Calls" --> Routes
    API -- "Middleware" --> Middleware
    DBLayer -- "SQL" --> PG
```

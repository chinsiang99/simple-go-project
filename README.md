# simple-go-project
Simple go project for me to learn about go

## Go Gin Backend Project Structure

```bash
myapp/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/              # HTTP handlers
│   │   │   ├── auth.go
│   │   │   ├── user.go
│   │   │   └── health.go
│   │   ├── middlewares/           # Custom middlewares
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   ├── logger.go
│   │   │   └── rate_limit.go
│   │   └── routes/                # Route definitions
│   │       └── routes.go
│   ├── config/                    # Configuration
│   │   └── config.go
│   ├── database/                  # Database related
│   │   ├── migrations/
│   │   │   ├── 001_create_users.sql
│   │   │   └── 002_create_posts.sql
│   │   └── database.go
│   ├── models/                    # GORM models
│   │   ├── user.go
│   │   ├── post.go
│   │   └── base.go
│   ├── repositories/              # Data access layer
│   │   ├── interfaces/
│   │   │   └── user_repository.go
│   │   └── user_repository.go
│   ├── services/                  # Business logic
│   │   ├── interfaces/
│   │   │   └── user_service.go
│   │   └── user_service.go
│   └── utils/                     # Utility functions
│       ├── logger/
│       │   └── logger.go
│       ├── validator/
│       │   └── validator.go
│       └── response/
│           └── response.go
├── pkg/                           # Public packages
│   └── errors/
│       └── errors.go
├── docs/                          # Swagger documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── tests/                         # Test files
│   ├── integration/
│   │   └── user_test.go
│   ├── unit/
│   │   ├── handlers/
│   │   │   └── user_test.go
│   │   ├── services/
│   │   │   └── user_test.go
│   │   └── repositories/
│   │       └── user_test.go
│   └── fixtures/
│       └── test_data.go
├── scripts/                       # Build and deployment scripts
│   ├── build.sh
│   └── migrate.sh
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env.example                   # Environment variables template
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
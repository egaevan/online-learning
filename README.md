# Online Learning

### Run Project

On the root directory, run this command:

```
go run app/main.go
```

## Directory structure

```
.
├── app                     # Main program of the app
├── config                  # Collection of config functions of the app (config reader, database connection, etc)
├── constant                # Collection of constants
├── delivery                # Delivery layer of the app
│   └── rest
│       ├── handler.go      # 
│       └── middleware.go   # 
├── go.mod                  # Go module file (collection of Go packages)
├── go.sum                  # Go sum file
├── model                   # Enterprise Business Logic and data structures
├── repository              # Repostiory layer of the app
└── usecase                 # Use case or business logic layer of the app
```

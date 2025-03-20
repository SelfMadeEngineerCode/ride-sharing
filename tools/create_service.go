package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	serviceName := flag.String("name", "", "Name of the service (e.g., user, payment)")
	flag.Parse()

	if *serviceName == "" {
		fmt.Println("Please provide a service name using -name flag")
		os.Exit(1)
	}

	// Create service directory structure
	basePath := filepath.Join("services", *serviceName+"-service")
	dirs := []string{
		"cmd",
		"internal/domain",
		"internal/service",
		"internal/infrastructure/events",
		"internal/infrastructure/grpc",
		"internal/infrastructure/repository",
		"pkg/types",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Create an empty README.md
	readmePath := filepath.Join(basePath, "README.md")
	readmeContent := fmt.Sprintf(`# %s service

This service handles all %s-related operations in the system.

## Architecture

The service follows Clean Architecture principles with the following structure:

` + "```" + `
services/%s-service/
├── cmd/                    # Application entry points
│   └── main.go            # Main application setup
├── internal/              # Private application code
│   ├── domain/           # Business domain models and interfaces
│   ├── service/          # Business logic implementation
│   │   └── service.go    # Service implementations
│   └── infrastructure/   # External dependencies implementations (abstractions)
│       ├── events/       # Event handling (RabbitMQ)
│       ├── grpc/         # gRPC server handlers
│       └── repository/   # Data persistence
├── pkg/                  # Public packages
│   └── types/           # Shared types and models
└── README.md            # This file
` + "```" + `

### Layer Responsibilities

1. **Domain Layer** (` + "`internal/domain/`" + `)
   - Contains business domain interfaces
   - Defines contracts for repositories and services
   - Pure business logic, no implementation details

2. **Service Layer** (` + "`internal/service/`" + `)
   - Implements business logic
   - Uses repository interfaces
   - Coordinates between different parts of the system

3. **Infrastructure Layer** (` + "`internal/infrastructure/`" + `)
   - ` + "`repository/`" + `: Implements data persistence
   - ` + "`events/`" + `: Handles event publishing and consuming
   - ` + "`grpc/`" + `: Handles gRPC communication

4. **Public Types** (` + "`pkg/types/`" + `)
   - Contains shared types and models
   - Can be imported by other services

## Key Benefits

1. **Dependency Inversion**: Services depend on interfaces, not implementations
2. **Separation of Concerns**: Each layer has a specific responsibility
3. **Testability**: Easy to mock dependencies for testing
4. **Maintainability**: Clear boundaries between components
5. **Flexibility**: Easy to swap implementations without affecting business logic
`, *serviceName, *serviceName, *serviceName)

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created %s service structure in %s\n", *serviceName, basePath)
	fmt.Println("\nDirectory structure created:")
	fmt.Printf(`
services/%s-service/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── domain/           # Business domain models and interfaces
│   │   └── %s.go         # Core domain interfaces
│   ├── service/          # Business logic implementation
│   │   └── service.go    # Service implementations
│   └── infrastructure/   # External dependencies implementations (abstractions)
│       ├── events/       # Event handling (RabbitMQ)
│       ├── grpc/         # gRPC server handlers
│       └── repository/   # Data persistence
├── pkg/                  # Public packages
│   └── types/           # Shared types and models
└── README.md            # This file
`, *serviceName, *serviceName)
} 
# Neura Go SDK

The official Go SDK for [Neura OS](https://neura-os.com), the Control Plane for Machine Intelligence.

## Installation

```bash
go get github.com/neura-os/neura-go
```

## Getting an API Key

**Option 1: Web Dashboard (Recommended)**

1. Visit [app.neura-os.com](https://app.neura-os.com).
2. Register for an account.
3. Generate an API Key from the dashboard.

**Option 2: Programmatic Registration**

You can also obtain an API key by registering your service using the SDK:

```go
package main

import (
	"fmt"
	"log"

	neura "github.com/neura-os/neura-go"
)

func main() {
	// Initialize client without an API Key for registration
	client, err := neura.NewClient(neura.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Register to get a new API Key
	resp, err := client.Auth.Register(neura.AuthRequest{
		OrgID: "my-org",
		Name:  "service-name",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("API Key: %s\n", resp.APIKey)
}
```

## Usage

### Initialization

```go
package main

import (
	"fmt"
	"log"
	
	neura "github.com/neura-os/neura-go"
)

func main() {
	// Initialize the client
	// API Key is automatically read from NEURA_API_KEY env var if not provided
	client, err := neura.NewClient(neura.Config{
		APIKey: "your-api-key", 
	})
	if err != nil {
		log.Fatal(err)
	}

	// Use the client...
}
```

### Making a Decision

```go
req := neura.DecisionRequest{
    Context: map[string]interface{}{
        "user_id": "123",
        "action":  "login",
    },
}

decision, err := client.Decide(req)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Outcome: %s\n", decision.Outcome)
```

## Features

- **Decision Engine**: Request decisions based on context.
- **Validation**: Validate outcomes.
- **Memory**: Store and retrieve contextual memory.
- **Authentication**: Manage identity registration.

## License

© 2026 NEURA OS. All rights reserved.

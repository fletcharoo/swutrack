# CLAUDE.md

## Project Context

### Project Overview
The name of this project is "swutrack". swutrack is a SaaS platform that helps competitive and casual players improve their gameplay through game result tracking and data-driven analytics. Players can store versioned decklists, log games with detailed opponent information, and access visual dashboards showing win rates, matchup analysis, and performance trends to make informed deck choices and tournament preparation decisions.

### Key Features
1. OAuth authentication
2. Deck management with versioning
3. Game tracking with relevant details (deck played, opponent's leader/base, result)
4. Statistics dashboard with charts and graphs

### Decision Making
When implementing features, prioritize:
1. User experience and core functionality
2. Data consistency and security
3. Performance and scalability
4. Code maintainability

## Development Environment

### Common Developer Commands
- `make test`: Run all tests
- `make dc/up`: Start local environment
- `make dc/down`: Stop local environment
- `make migration/add`: Create a new database migration

### Technical Stack
- **Backend:** Go with REST API
- **Database:** Postgres with Goose migrations
- **Frontend:** htmx + server-rendered templates
- **Auth:** OAuth
- **Storage:** S3-compatible buckets
- **Testing:** testify/assert and testify/require

### Local Development Environment
- Local setup uses Docker Compose
- Database runs on port 5432
- Application runs on port 8080

### Important Files
- **`service/main.go`**: Application entry point that initializes and starts the HTTP server
- **`service/transport/httpapi/server.go`**: HTTP server implementation with configuration and lifecycle management
- **`service/transport/httpapi/routes.go`**: Route configuration that maps endpoints to handlers
- **`service/transport/httpapi/handlers.go`**: HTTP request handlers implementation
- **`go.mod`**: Go module definition file for dependency management
- **`Makefile`**: Build automation and development commands
- **`compose.yaml`**: Docker Compose configuration for local development environment
- **`service/Dockerfile`**: Multi-stage Docker build configuration for the application
- **`CLAUDE.md`**: Project instructions and coding standards for AI assistance
- **`docs/api-endpoints.md`**: Comprehensive API endpoint documentation
- **`docs/db-schema.sql`**: Complete database schema design
- **`docs/star-wars-unlimited.md`**: Domain knowledge documentation for game mechanics

### Project Structure
```
swutrack/
├── service/                     # Go application code
│   ├── transport/              # Transport layer
│   │   └── httpapi/           # HTTP API implementation
│   │       ├── handlers.go    # Request handlers
│   │       ├── routes.go      # Route definitions
│   │       └── server.go      # HTTP server setup
│   ├── main.go                # Application entry point
│   └── Dockerfile             # Container build configuration
├── docs/                       # Project documentation
│   ├── api-endpoints.md       # API specification
│   ├── db-schema.sql          # Database design
│   └── star-wars-unlimited.md # Domain knowledge
├── .claude/                    # Claude AI configuration
│   ├── commands/              # Custom AI commands
│   │   ├── fcc:commit.md
│   │   ├── fcc:create-prd.md
│   │   ├── fcc:plan.md
│   │   ├── fcc:review.md
│   │   ├── fcc:update-claudemd.md
│   │   └── fcc:update-documentation.md
│   └── settings.local.json    # Local Claude settings
├── temp/                       # Temporary files (git ignored)
├── go.mod                      # Go module definition
├── Makefile                    # Build automation
├── compose.yaml                # Docker Compose setup
├── README.md                   # Project overview
├── LICENSE                     # MIT License
├── CLAUDE.md                   # Project instructions for AI
└── .gitignore                  # Git ignore rules
```

### Documentation
- **Database schema:** `docs/db-schema.sql`
- **API endpoints:** `docs/api-endpoints.md`
- **Domain knowledge:** `docs/star-wars-unlimited.md`

## Technical Standards

### Database
- Use Goose for migrations in `migrations/` directory
- Migration naming: `YYYYMMDDHHMMSS_description.sql`
- Always include both up and down migrations

### HTTP/API Conventions
- RESTful endpoints following `/api/v1/resource` pattern
- Use appropriate HTTP status codes
- JSON request/response format
- Include request validation middleware

### Authentication
- OAuth 2.0 flow with Google/Discord providers
- JWT tokens for session management
- Middleware for protected routes

### Error Handling
- Use wrapped errors with context: `fmt.Errorf("failed to get user %q: %w", userID, err)`
- Log errors at the boundary where they're handled
- Return structured errors for API responses

### Logging
- Use structured logging with context
- Log at appropriate levels (debug, info, warn, error)
- Include relevant request IDs and user context

### Configuration
- Use environment variables for all configuration
- Provide sensible defaults where possible
- Document all required environment variables

### Timestamps
- All timestamps must be in UTC timezone

## Code Quality

### Code Style & Conventions
In addition to using idiomatic Go conventions, all code you write **MUST** align with the following style and conventions:
- Where possible, use already established patterns
- All names should be descriptive and avoid using contractions
- Follow early return principles
- All functions **MUST** assert their input and gracefully handle invalid inputs (e.g., strings aren't empty, pointers aren't nil)
- All functions **MUST** have named return parameters (e.g., `(err error)`)
- All code **MUST** be documented with comments that align with godoc standards
- Test assertions **MUST** be done with the `stretchr/testify/assert` package
- Test requirements **MUST** be done with the `stretchr/testify/require` package
- Table tests **MUST** align with the template described in the `Examples` section
- All code **MUST** align with the single responsibility principle
- All code **MUST** align with separation of concerns
- Only use pointers when data needs to be mutated or when dealing with large structs where copying would be inefficient. Return values by value when possible
- Always put an empty newline a the end of files
- **ONLY** use an empty return if it's an error return. If you're returning data or a nil error, make that explicit through the return.

### Clarifying Questions
When you ask the user clarifying questions, you **MUST** follow the following rules:
- The goal is to understand the "what" and "why" of the feature, not necessarily the "how" (which the developer will figure out). Make sure to provide options in numbered lists so I can respond easily with my selections
- Only ask **ONE** clarifying question at a time
- Use the responses from clarifying questions to improve all tasks you perform

## References

### Function Error Handling Example
```go
// GetUserByID retrieves a user by their unique identifier.
func GetUserByID(ctx context.Context, userID string) (user User, err error) {
    if userID == "" {
        err = fmt.Errorf("user ID cannot be empty")
        return
    }
}
```

### Table Test Template
```go
func Test_[Insert test name](t *testing.T) {
    testCases := map[string]struct{
        // inputs and expectations
    }{
        // test cases
    }

    for name, tc := range testCases {
        t.Run(name, func(t *testing.T) {
            // run the test and check expectations
        })
    }
}
```

### Clarifying Questions Examples
The AI **MUST** adapt its questions based on the task and requirements, but here are some common areas to explore:
*   **Problem/Goal:** "What problem does this feature solve for the user?" or "What is the main goal we want to achieve with this feature?"
*   **Target User:** "Who is the primary user of this feature?"
*   **Core Functionality:** "Can you describe the key actions a user should be able to perform with this feature?"
*   **User Stories:** "Could you provide a few user stories? (e.g., As a [type of user], I want to [perform an action] so that [benefit].)"
*   **Acceptance Criteria:** "How will we know when this feature is successfully implemented? What are the key success criteria?"
*   **Scope/Boundaries:** "Are there any specific things this feature *should not* do (non-goals)?"
*   **Data Requirements:** "What kind of data does this feature need to display or manipulate?"
*   **Design/UI:** "Are there any existing design mockups or UI guidelines to follow?" or "Can you describe the desired look and feel?"
*   **Edge Cases:** "Are there any potential edge cases or error conditions we should consider?"

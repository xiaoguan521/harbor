# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

Harbor is an open source cloud native registry that stores, signs, and scans container images and Helm charts. It's a CNCF (Cloud Native Computing Foundation) project built primarily in Go with an Angular web UI.

## Common Development Commands

### Build Commands

```bash
# Full build including compile, build images, and prepare
make all

# Compile Go binaries only
make compile

# Build Docker images
make build

# Install Harbor (compile + build + prepare + start)
make install
```

### Test Commands

```bash
# Run Go tests in specific package
cd src/[package]
go test -v ./...

# Run UI tests
cd src/portal/lib
npm run test

# Run linting and code checks
make go_check

# Run specific linters
make lint           # Go code linting
make commentfmt     # Comment format checking
make misspell       # Spell checking
```

### Development Server Commands

```bash
# Start Harbor services
make start

# Stop Harbor services  
make down

# Restart Harbor
make restart

# Generate API code from swagger
make gen_apis

# Generate mocks for testing
make gen_mocks
```

### UI Development Commands

```bash
cd src/portal

# Install dependencies
npm install

# Copy proxy config template
cp proxy.config.mjs.temp proxy.config.mjs

# Start development server
npm run start
# Access at https://localhost:4200
```

### Package and Deploy Commands

```bash
# Create offline installer package
make package_offline

# Create online installer package  
make package_online

# Push images to registry
make pushimage -e REGISTRYSERVER=<server> REGISTRYUSER=<user> REGISTRYPASSWORD=<pass>
```

### Cleanup Commands

```bash
# Clean all build artifacts
make cleanall

# Clean specific components
make cleanbinary      # Remove compiled binaries
make cleanimage       # Remove Docker images
make cleanpackage     # Remove installer packages
```

## Architecture Overview

Harbor follows a microservices architecture with the following core components:

### Core Services
- **harbor-core**: Main API server and business logic, handles authentication, project management, repository management
- **harbor-jobservice**: Background job processing service for replication, garbage collection, vulnerability scanning
- **harbor-portal**: Angular-based web UI
- **harbor-db**: PostgreSQL database for metadata storage
- **harbor-redis**: Redis for job queue and caching
- **harbor-registry**: Docker registry based on Docker Distribution
- **harbor-registryctl**: Controller for managing registry storage and operations

### Optional Components
- **harbor-trivy-adapter**: Trivy vulnerability scanner adapter
- **harbor-exporter**: Prometheus metrics exporter
- **harbor-log**: Centralized logging service

### Key Source Code Structure

```
src/
├── core/           # Main Harbor API server and business logic
├── jobservice/     # Background job processing service
├── portal/         # Angular web UI application
├── registryctl/    # Registry controller service
├── common/         # Shared utilities (dao, models, config, etc.)
├── controller/     # Business logic controllers
├── pkg/           # Reusable packages and utilities
├── server/        # API server implementation
└── cmd/           # Command-line tools and utilities
```

### Programming Model
Harbor uses a layered architecture:
- **Controllers**: Business logic layer
- **Managers**: Service management layer  
- **DAOs**: Data access objects for database operations

### API Design
- REST APIs generated from OpenAPI 2.0 specification (`api/v2.0/swagger.yaml`)
- Uses go-swagger for API code generation
- API handlers located in `src/server/v2.0/handler/`

## Build System Details

### Go Build Configuration
- **Go Version**: 1.24.6 (as specified in `src/go.mod`)
- **Build Tools**: Uses Docker containers for consistent builds
- **Go Build Image**: `golang:1.24.6`
- **Build Flags**: Includes version info via ldflags (`-X pkg/version.GitCommit` and `-X pkg/version.ReleaseVersion`)

### Docker Build Process
- **Base Images**: Uses photon OS base images
- **Multi-stage Builds**: Separates build and runtime environments
- **Image Namespace**: `goharbor/` (configurable via `IMAGENAMESPACE`)

### Frontend Build
- **Framework**: Angular with Clarity Design System
- **Node Version**: Defined in `.nvmrc` file
- **Build Tool**: npm/Angular CLI
- **Development Server**: Runs on https://localhost:4200

## Testing Strategy

### Go Testing
- Unit tests using Go's built-in testing framework
- Mock generation via mockery (config in `src/.mockery.yaml`)
- Use testify/mock for controller and manager testing
- Coverage reporting integrated with CI

### Frontend Testing  
- Jasmine and Karma test framework
- Angular testing utilities
- Component and integration tests

### Linting and Code Quality
- golangci-lint configuration in `src/.golangci.yaml`
- Enabled linters: bodyclose, errcheck, goheader, govet, ineffassign, misspell, revive, staticcheck, whitespace
- Code formatting with gofmt and goimports
- Comment format checking and spell checking

## Development Workflow Notes

### API Changes
1. Update `api/v2.0/swagger.yaml`
2. Run `make gen_apis` to regenerate server code
3. Implement handlers in `src/server/v2.0/handler/`

### Mock Generation
1. Add interface to `src/.mockery.yaml` config
2. Run `make gen_mocks` to generate mocks
3. Use generated mocks in tests

### Database Migrations
- Migration scripts in `src/migration/`
- Standalone migrator tool for DB upgrades

### Dependency Management
- Go modules with `src/go.mod`
- Custom module replacements for Harbor-specific forks
- NPM for frontend dependencies

## Environment Variables and Configuration

- Configuration template: `make/harbor.yml.tmpl`
- Harbor uses Beego framework configuration patterns
- Environment-specific settings handled via docker-compose

## Version Information

Current version: v2.15.0 (from VERSION file)

## CI/CD Integration

- GitHub Actions workflows in `.github/workflows/`
- Docker image builds and pushes
- Automated testing and linting
- Security scanning integration
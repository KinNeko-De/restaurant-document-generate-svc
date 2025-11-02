# GitHub Copilot Instructions

This file contains specific instructions and reminders for GitHub Copilot when working on this repository.

## Dependency Updates

### OpenTelemetry Dependencies

When updating OpenTelemetry dependencies (go.opentelemetry.io/otel/*), you must also update the semantic conventions import in the following file:

**File:** `internal/app/operation/metric/metric.go`
**Import to update:** 
```go
semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
```

**Action Required:**
- Update the version number in the semconv import to match the OpenTelemetry version being used
- The semconv version should typically match or be compatible with the main otel version. Try the exact same version first, and if there are issues, check the OpenTelemetry release notes for the correct semconv version.
- Example: If updating to otel v1.38.0, update semconv import to v1.38.0

**Why this is needed:**
The semantic conventions package version needs to stay in sync with the main OpenTelemetry packages to avoid schema URL conflicts and ensure compatibility.

## Go Version Updates

When updating the Go version:
1. Update `go.mod` file
2. Update `.github/workflows/ci.yml` - change the `go-version` field
3. Run `go get -u ./...` to update all dependencies
4. Run `go mod tidy` to clean up dependencies
5. Verify build and core tests still pass

## Code Quality

- Maintain existing code style and patterns
- Ensure all exported functions have proper documentation
- Add unit tests for new functionality
- Keep dependency updates atomic and well-documented

## Testing
- All tests should always pass
- Run `go test ./internal/app/document ./internal/app/server` to verify core functionality
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

**Step-by-step process (minimize premium requests):**
1. Update the import version in `metric.go` to match the new OpenTelemetry version
2. Add the semconv as a direct dependency: `go get go.opentelemetry.io/otel/semconv/v1.X.0` (replace X with version)
3. If step 2 fails (package doesn't exist), try one version lower until it succeeds
4. Run `go mod tidy` once to clean up
5. Verify with: `go test ./internal/app/operation/metric`

**Version compatibility notes:**
- The semconv version should match the main otel version when possible
- If exact version doesn't exist, use the highest available version (typically N-1)
- Example: For otel v1.38.0, try v1.38.0 first, then v1.37.0 if needed
- The semconv package is managed separately and may lag behind main otel releases

**Why this is needed:**
The semantic conventions package version needs to stay in sync with the main OpenTelemetry packages to avoid schema URL conflicts and ensure compatibility.

## Go Version Updates

When updating the Go version:
1. Update `go.mod` file
2. Update `.github/workflows/ci.yml` - change the `go-version` field
3. Run `go get -u ./...` to update all dependencies
4. Run `go mod tidy` to clean up dependencies
5. Verify build and core tests still pass

## GitHub Actions Workflow Updates

Always use the newest stable versions of GitHub Actions in workflows. When updating action versions, follow the step-by-step process below to avoid introducing CI breakage and to minimize unnecessary network requests or trial-and-error steps.

Step-by-step process for updating action versions (minimize premium/trial requests):
1. Search the workflow files under `.github/workflows/` for `uses:` entries that reference third-party actions.
2. For each action, visit the action's official repository or the Marketplace page to find the latest stable release tag. Prefer a pinned major version (for example `actions/checkout@v4`). Also bump the major version when a new major release is available. In most cases this is safe and we can roll it back if the pipeline breaks. A successful pipeline run is needed to complete a pull request. You can do network calls for that.
3. Update the workflow to the new tag. Example replacements:
	- `actions/checkout@v3` → `actions/checkout@v4`
	- `actions/setup-go@v4` → `actions/setup-go@v5`
	- `docker/build-push-action@v4` → `docker/build-push-action@v5`

Best practices and safety notes:
- Prefer pinned version tags (e.g., `@v4`) rather than floating tags like `@main` or `@master`.
- Read the action's changelog for breaking changes before updating.

Examples of common actions in this repo (verify these when updating):
- `actions/checkout@v4`
- `actions/setup-go@v5`
- `codecov/codecov-action@v4`
- `docker/build-push-action@v5`
- `actions/github-script@v7`

Rollback guidance:
- If an updated action breaks CI in a way that's not immediately fixable, revert the workflow commit

## Code Quality

- Maintain existing code style and patterns
- Ensure all exported functions have proper documentation
- Add unit tests for new functionality
- Keep dependency updates atomic and well-documented

## Testing
- All tests should always pass
- Run `go test ./internal/app/document ./internal/app/server` to verify core functionality
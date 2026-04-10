# Go submodule workflow

Use this reference when the task edits SDK source, tests, snippets, module metadata, or generated files inside `submodules/go`.

## Layout

- `barcode`: generated public client code, request/response models, enums, auth context keys, and `barcode/jwt`.
- `docs`: generated Markdown API and model docs.
- `examples`: small standalone sample apps for token fetch, generate, and scan flows.
- `snippets`: repo documentation snippets grouped by generate and read scenarios.
- `test`: Go test coverage for generate, recognize, scan, configuration, JWT auth, and API error cases.
- `testdata`: sample images used by tests, snippets, and examples.
- `scripts`: build, lint, format, test, snippet-runner, and post-processing helpers.
- `go.mod` and `go.sum`: module metadata for `github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4`.
- `README.md`: end-user setup and usage overview.

## Validation

On Windows, run repo scripts and Make targets through WSL.

From `submodules/go`:

- `make build`
- `make test`
- `make lint`
- `make format`

Useful composite targets:

- `make release`: runs format, lint, dependency cleanup, build, and test.
- `make after-gen`: runs init, formatting, example insertion, and module cleanup after regeneration.

`make test` does more than plain `go test`:

- `scripts/test.sh` runs `go test -v $(go list ./... | grep -v snippets)` so snippet packages are excluded from the normal test command.
- `scripts/run_snippets.sh` creates a temporary `snippets_test` folder and executes every snippet individually through `scripts/run_snippet.sh`.
- `scripts/run_snippet.sh` copies credentials into a runnable temp file with `scripts/insert-credentials.py`, then runs `go run` on that snippet.

Treat snippet failures as consumer-facing regressions, not just sample breakage.

## Test configuration

- Template values live in `test/configuration.example.json`.
- Tests read `test/configuration.json` first.
- If `test/configuration.json` is absent, `test/test_config.go` falls back to environment variables using the `TEST_` prefix.
- Common environment variables include `TEST_JWT_CLIENT_ID`, `TEST_JWT_CLIENT_SECRET`, `TEST_JWT_TOKEN_URL`, `TEST_JWT_ACCESS_TOKEN`, and `TEST_API_BASE_PATH`.

`test/helpers.go` is the central test harness for constructing the client and auth context.

## Regenerated code workflow

If you change generated SDK code in this mono-repo:

1. Make the desired SDK edit in `submodules/go` so the target behavior is clear.
2. Mirror the change in the matching template under `codegen/Templates` when the file is generated.
3. Stage the Go submodule changes.
4. From the repo root, run `make go`.
5. Ensure `submodules/go` has no new unstaged diffs after regeneration.
6. If regeneration reintroduces old code, keep fixing templates until the generated output matches the intended SDK change.

## Useful anchors

- `barcode/api_generate.go`: generate endpoint shapes and `Generate*` method signatures.
- `barcode/api_recognize.go`: recognize endpoint shapes and `Recognize*` method signatures.
- `barcode/api_scan.go`: scan endpoint shapes and `Scan*` method signatures.
- `barcode/configuration.go`: base path defaults, default headers, and user-agent setup.
- `barcode/jwt/jwt.go`: `jwt.Config`, token fetching, and pre-fetched access-token behavior.
- `test/helpers.go`: how repo tests construct the auth context and API client.
- `test/api_generate_test.go`, `test/api_recognize_test.go`, `test/api_scan_test.go`: end-to-end API usage patterns.

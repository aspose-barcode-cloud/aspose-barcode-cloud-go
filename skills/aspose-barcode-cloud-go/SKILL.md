---
name: aspose-barcode-cloud-go
description: "Write or update Go code that uses the Aspose.BarCode Cloud SDK for Go (module `github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4`) to generate, recognize, or scan barcodes through Aspose's cloud REST API. Use this skill whenever the user wants barcode work in Go, touches files under `submodules/go`, or mentions `GenerateAPI`, `RecognizeAPI`, `ScanAPI`, `GenerateParams`, `RecognizeBase64Request`, `ScanBase64Request`, `barcode.ContextJWT`, or `github.com/antihax/optional`. The Go SDK has several easy-to-miss idioms: the `/v4` import-path suffix, auth flowing through `context.Context` with `barcode.ContextJWT`, `optional.New*` wrappers only on `*Opts` structs, `GenerateBody` vs `RecognizeBase64` and `ScanBase64` naming, and GET recognize/scan methods requiring a public `fileUrl`, so consult this skill instead of guessing."
---

# Aspose.BarCode Cloud SDK for Go

The Go SDK is a thin generated client over the Aspose BarCode Cloud REST API. Most tasks come down to choosing the right API service (`GenerateAPI`, `RecognizeAPI`, or `ScanAPI`), choosing the right transport variant (GET, body/base64, or multipart), and wiring authentication through `context.Context` correctly.

The module path includes a major-version suffix. Install and import `github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4`, not the unsuffixed path.

## Quick start

Use these imports in most Go examples:

```go
import (
    "context"

    "github.com/antihax/optional"

    "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
    "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
)
```

Prefer one shared API client:

```go
jwtConf := jwt.NewConfig(clientID, clientSecret)
authCtx := context.WithValue(
    context.Background(),
    barcode.ContextJWT,
    jwtConf.TokenSource(context.Background()),
)

client := barcode.NewAPIClient(barcode.NewConfiguration())
```

If the task is repo maintenance inside `submodules/go`, read `references/repo-workflow.md`. If the task needs the closest existing snippet, example, or test, read `references/snippet-map.md`.

## Authentication

Use one of these two patterns:

1. Let the SDK fetch and refresh JWT tokens for you through `barcode.ContextJWT`.

```go
jwtConf := jwt.NewConfig(clientID, clientSecret)
authCtx := context.WithValue(
    context.Background(),
    barcode.ContextJWT,
    jwtConf.TokenSource(context.Background()),
)

client := barcode.NewAPIClient(barcode.NewConfiguration())
```

2. Inject a pre-fetched bearer token when snippets or CI already provide one.

```go
config := barcode.NewConfiguration()
config.AddDefaultHeader("Authorization", "Bearer "+token)

client := barcode.NewAPIClient(config)
authCtx := context.Background()
```

Inside this repo, snippet files often check `TEST_JWT_ACCESS_TOKEN` first and fall back to client credentials. Tests load `test/configuration.json` first, then build the same auth flow from `TEST_*` environment variables when the file is absent.

## Choose the right API shape

Pick the operation first:

- `GenerateAPI`: create a barcode image.
- `RecognizeAPI`: decode one or more expected barcode types and optionally tune recognition quality.
- `ScanAPI`: auto-detect barcode types with the smallest API surface.

Then pick the transport variant based on what the caller has:

- Public image URL: use `RecognizeAPI.Recognize` or `ScanAPI.Scan`. `fileUrl` must be a public URL, not a local path.
- Local file on disk: use `RecognizeMultipart` or `ScanMultipart`.
- Raw bytes already in memory: base64-encode them yourself and use `RecognizeBase64` or `ScanBase64`.
- Short text plus query-style parameters for generation: use `Generate`.
- Structured generate payload or longer data: use `GenerateBody`.
- Multipart-form generation: use `GenerateMultipart` only when the caller explicitly needs that shape.

Key method names:

- `Generate`
- `GenerateBody`
- `GenerateMultipart`
- `Recognize`
- `RecognizeBase64`
- `RecognizeMultipart`
- `Scan`
- `ScanBase64`
- `ScanMultipart`

## Non-obvious SDK rules

1. The `/v4` suffix is mandatory in every import path. Omitting it breaks builds.
2. Authentication usually flows through `context.Context` with `barcode.ContextJWT`, not through a special auth client type.
3. `Generate` returns `[]byte`. Save the result with `os.WriteFile` or another writer.
4. `RecognizeBase64` and `ScanBase64` expect a base64 string in the request model. The SDK does not encode raw bytes for you.
5. `RecognizeBase64Request.BarcodeTypes` is a slice of `DecodeBarcodeType`, but `Recognize` and `RecognizeMultipart` take a single `DecodeBarcodeType`.
6. `ScanAPI` does not take barcode types or recognition tuning knobs. Use it when the caller wants auto-detection.
7. Optional-parameter structs (`*Opts`) use `optional.New*` wrappers. Body structs such as `GenerateParams`, `RecognizeBase64Request`, and `ScanBase64Request` use plain values.
8. GET-based recognize and scan methods only work with remote files reachable by URL. For local files, use multipart or base64.
9. Every recognize or scan result is a `BarcodeResponseList`. Iterate `result.Barcodes` and read `BarcodeValue`, `Type`, `Region`, and `Checksum`.
10. API failures may come back as `barcode.GenericAPIError`; keep the returned `*http.Response` around when debugging status codes and response bodies.

## Common patterns

Generate and save a QR code:

```go
opts := &barcode.GenerateAPIGenerateOpts{
    ImageFormat:  optional.NewInterface(barcode.BarcodeImageFormatPng),
    TextLocation: optional.NewInterface(barcode.CodeLocationNone),
}

imageBytes, _, err := client.GenerateAPI.Generate(
    authCtx,
    barcode.EncodeBarcodeTypeQR,
    "hello from Go",
    opts,
)
if err != nil {
    return err
}

return os.WriteFile("qr.png", imageBytes, 0644)
```

Recognize specific barcode types from bytes already in memory:

```go
req := barcode.RecognizeBase64Request{
    BarcodeTypes: []barcode.DecodeBarcodeType{
        barcode.DecodeBarcodeTypeQR,
        barcode.DecodeBarcodeTypeCode128,
    },
    FileBase64: base64.StdEncoding.EncodeToString(imageBytes),
}

result, _, err := client.RecognizeAPI.RecognizeBase64(authCtx, req)
```

Auto-scan a local file without specifying the barcode type:

```go
file, err := os.Open("unknown.png")
if err != nil {
    return err
}
defer file.Close()

result, _, err := client.ScanAPI.ScanMultipart(authCtx, file)
```

## Working in this repo

Read `references/repo-workflow.md` when the task changes SDK source, tests, snippets, module metadata, or generated code in `submodules/go`.

Read `references/snippet-map.md` when the task needs the closest existing pattern for generate, recognize, scan, auth, or repo-test scenarios.

## Final checklist

1. Use the correct module path: `github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/...`.
2. Choose `GenerateAPI`, `RecognizeAPI`, or `ScanAPI` based on whether the caller is generating, recognizing known types, or auto-scanning unknown types.
3. Use `barcode.ContextJWT` for the normal auth flow and keep pre-fetched-token handling consistent with the surrounding repo code.
4. Pick GET only for public URLs, multipart for local files, and base64 request models for in-memory bytes.
5. Use `optional.New*` only on `*Opts` structs, not on request-body structs.
6. Treat generate responses as image bytes and recognize/scan responses as `result.Barcodes`.
7. When changing the repo, validate with the submodule workflow in `references/repo-workflow.md`.

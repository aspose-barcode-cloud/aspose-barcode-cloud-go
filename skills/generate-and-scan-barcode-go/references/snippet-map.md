# Snippet and example map

Use this reference when you want the closest existing pattern before writing new Go SDK code or when updating docs, snippets, and examples.

## Small end-user examples

- `examples/generate/example.go`: minimal QR generation to a local file.
- `examples/scan/example.go`: minimal auto-scan from a local file with `ScanMultipart`.
- `examples/fetch_token/example.go`: manual token fetch through `jwt.Config`.
- `snippets/manual_fetch_token.go`: raw HTTP client-credentials token fetch without the SDK's auth helper.

## Generate patterns

- `snippets/generate/save/generate_get.go`: simple `Generate` and save-to-file flow.
- `snippets/generate/save/generate_body.go`: `GenerateBody` with `GenerateParams`.
- `snippets/generate/save/generate_multipart.go`: multipart generation flow.
- `snippets/generate/set_text/*`: `EncodeData` and `EncodeDataType` examples.
- `snippets/generate/set_size/*`: width, height, resolution, and units examples.
- `snippets/generate/set_colorscheme/*`: foreground and background color examples.
- `snippets/generate/appearance/*`: richer `BarcodeImageParams` examples across GET, body, and multipart variants.

## Recognize and scan patterns

- `snippets/read/set_source/recognize_get.go`: recognize from a public URL.
- `snippets/read/set_source/recognize_multipart.go`: recognize from a local file.
- `snippets/read/set_source/recognize_body.go`: recognize from base64 bytes.
- `snippets/read/set_source/scan_get.go`: auto-scan from a public URL.
- `snippets/read/set_source/scan_multipart.go`: auto-scan from a local file.
- `snippets/read/set_source/scan_body.go`: auto-scan from base64 bytes.
- `snippets/read/set_target_types/*`: choosing `DecodeBarcodeType` vs `[]DecodeBarcodeType`.
- `snippets/read/set_quality/*`: `RecognitionMode` examples.
- `snippets/read/set_image_kind/*`: `RecognitionImageKind` examples.

## Tests worth copying

- `test/api_generate_test.go`: generate via GET, body, and multipart variants.
- `test/api_recognize_test.go`: recognize via URL, base64 body, and multipart.
- `test/api_scan_test.go`: scan via URL, base64 body, and multipart.
- `test/jwt_test.go`: token-source and validation behavior.
- `test/barcode_error_test.go`: expected API error behavior.
- `test/configuration_test.go`: configuration defaults and header behavior.
- `test/test_config_test.go`: file-vs-env test configuration loading.

## Repo-specific auth anchors

- Most snippets check `TEST_JWT_ACCESS_TOKEN` first, then fall back to `jwt.NewConfig(...)`.
- `test/helpers.go` and `test/test_config.go` show the repo's preferred auth-context construction for tests.
- `barcode/jwt/jwt.go` is the source of truth for how `AccessToken`, `ClientID`, `ClientSecret`, and `TokenURL` interact.

module github.com/aspose-barcode-cloud/aspose-barcode-cloud-go

go 1.11

require (
	github.com/antihax/optional v1.0.0
	github.com/google/uuid v1.3.0
	// Do not upgrade Testify to >=1.7.0. It breaks Go 1.11 compatibility: "module requires Go 1.13"
	github.com/stretchr/testify v1.6.1
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
)

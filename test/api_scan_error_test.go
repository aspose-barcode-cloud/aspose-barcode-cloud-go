package test

import (
	"os"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

func TestWrongScanUrl(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.ScanAPI.Scan(authCtx, "not-a-url")

	requireAPIError(t, err)
}

func TestWrongScanBase64Content(t *testing.T) {
	client, authCtx := setup(t)

	_, _, err := client.ScanAPI.ScanBase64(authCtx, barcode.ScanBase64Request{
		FileBase64: "not-base64",
	})

	requireAPIError(t, err)
}

func TestWrongScanMultipartContent(t *testing.T) {
	client, authCtx := setup(t)

	file, err := os.CreateTemp(t.TempDir(), "not-image-*.txt")
	require.NoError(t, err)
	_, err = file.WriteString("not an image")
	require.NoError(t, err)
	_, err = file.Seek(0, 0)
	require.NoError(t, err)
	defer file.Close()

	_, _, err = client.ScanAPI.ScanMultipart(authCtx, file)

	requireAPIError(t, err)
}

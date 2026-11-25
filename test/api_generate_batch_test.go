package test

import (
	"archive/zip"
	"bytes"
	"testing"

	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/stretchr/testify/require"
)

func TestGenerateBatch(t *testing.T) {
	for _, format := range []barcode.BatchArchiveFormat{barcode.BatchArchiveFormatZip, barcode.BatchArchiveFormatTar, barcode.BatchArchiveFormatTarGzip} {
		t.Run(string(format), func(t *testing.T) {
			client, ctx := setup(t)
			data, _, err := client.GenerateAPI.GenerateBatch(ctx, barcode.BatchGenerateRequest{
				ArchiveFormat: format,
				Items: []barcode.BatchGenerateItem{{FileName: "qr.png", Generation: barcode.GenerateParams{
					BarcodeType: barcode.EncodeBarcodeTypeQR, EncodeData: barcode.EncodeData{Data: "Batch example"},
				}}},
			})
			require.NoError(t, err)
			require.NotEmpty(t, data)
			switch format {
			case barcode.BatchArchiveFormatZip:
				archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
				require.NoError(t, err)
				require.Len(t, archive.File, 1)
				require.Equal(t, "qr.png", archive.File[0].Name)
			case barcode.BatchArchiveFormatTarGzip:
				require.True(t, bytes.HasPrefix(data, []byte{0x1f, 0x8b}))
			default:
				require.Greater(t, len(data), 262)
				require.Equal(t, "ustar", string(data[257:262]))
			}
		})
	}
}

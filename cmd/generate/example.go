package main

import (
	"context"
	"fmt"
	"os"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/aspose_barcode_cloud"
)

func main() {
	fileName := "testdata/generated.png"

	//TODO: Move to common configuration
	cfg := api.NewConfiguration()
	cfg.BasePath = "http://localhost:47972/v3.0"
	client := api.NewAPIClient(cfg)
	ctx := context.WithValue(context.TODO(), api.ContextAccessToken, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1OTEyNzc1MDMsImV4cCI6MTU5MTM2MzkwMywiaXNzIjoiaHR0cHM6Ly9hcGktcWEuYXNwb3NlLmNsb3VkIiwiYXVkIjpbImh0dHBzOi8vYXBpLXFhLmFzcG9zZS5jbG91ZC9yZXNvdXJjZXMiLCJhcGkucGxhdGZvcm0iLCJhcGkucHJvZHVjdHMiXSwiY2xpZW50X2lkIjoiYmFyY29kZS5jbG91ZCIsImNsaWVudF91c2FnZSI6InBvc3QiLCJjbGllbnRfc3RvcmFnZSI6ImdldCIsImNsaWVudF9yZXN0cmljdGlvbnMiOiJnZXQiLCJjbGllbnRfZGVmYXVsdF9zdG9yYWdlIjoiMEZCOUI2MjctNzY5QS00MjJBLUJBMkUtODhBMUUzNEFDODA1IiwiY2xpZW50X2lkU3J2SWQiOiIxODkyMDEiLCJzY29wZSI6WyJhcGkucGxhdGZvcm0iLCJhcGkucHJvZHVjdHMiXX0.ZXYVEiZncgoEVHaFFg0p4MGGBlak2Z7OKzXhYu6WukEyYa9fCXw6vAoLUJqsam2-FKpXhZtssLBzyPrRP7kqtJ1gTNEMsSEDbK2HwLr6WzXfCn81h7VJiV6IOsqn0KY0K8Yb-Ge27G_325hJKzPp712hN6Cl1bTjG9lnfnPdZ2g1SGYMcQ_6062hTukXHDYTYY_1yr-SM_TfWOzIuQ8z1Wop_lmX2ur3hRPP1OyS33hwbxGbrErzXvlBFPxf5GWCSzZCaBk-JQYMRMH0QKUE-RPXtrwCwsd6ejvGN74JQZXpt__hUjGh7FaKmaOFfSL4H4VnMzCzbvmE9VH7D-Ol4g")

	data, resp, err := client.BarcodeApi.GetBarcodeGenerate(ctx, string(api.CODE128_EncodeBarcodeType), "Go SDK", nil)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(data)
	}
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	written, err := out.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d bytes to file %s\n", written, fileName)
}

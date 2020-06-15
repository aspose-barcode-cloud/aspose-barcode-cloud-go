package test

import (
	"encoding/json"
	"io/ioutil"
	"os"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
)

//Config is Configuration
type Config struct {
	JwtConfig jwt.Config        `json:"jwt"`
	APIConfig api.Configuration `json:"api"`
}

//NewConfig creates new Config from JSON-file
func NewConfig(fileName string) (*Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config, err := NewConfigFromJSON(bytes)
	if err != nil {
		return nil, err
	}

	return config, nil
}

//NewConfigFromJSON creates new Config from JSON-bytes
func NewConfigFromJSON(bytes []byte) (*Config, error) {
	config := Config{
		JwtConfig: *jwt.NewConfig("", ""),
		APIConfig: *api.NewConfiguration(),
	}
	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

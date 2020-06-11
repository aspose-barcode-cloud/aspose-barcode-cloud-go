package test

import (
	"encoding/json"
	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
	"io/ioutil"
	"os"
)

type Config struct {
	JwtConfig jwt.Config        `json:"jwt"`
	ApiConfig api.Configuration `json:"api"`
}

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

	config, err := NewConfigFromJson(bytes)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewConfigFromJson(bytes []byte) (*Config, error) {
	config := Config{
		JwtConfig: *jwt.NewConfig("", ""),
		ApiConfig: *api.NewConfiguration(),
	}
	err := json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

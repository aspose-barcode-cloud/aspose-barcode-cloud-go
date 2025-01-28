package test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"unicode"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/v4/barcode/jwt"
)

func camelToUpperSnake(str string) string {
	var builder strings.Builder

	prevCharWasLower := false
	for _, rn := range str {
		if unicode.IsUpper(rn) {
			if prevCharWasLower {
				builder.WriteRune('_')
			}
			builder.WriteRune(rn)
		} else {
			builder.WriteRune(unicode.ToUpper(rn))
		}

		prevCharWasLower = unicode.IsLower(rn)
	}

	return builder.String()
}

// Config is Configuration
type Config struct {
	JwtConfig jwt.Config        `json:"jwt"`
	APIConfig api.Configuration `json:"api"`
}

// NewTestConfig creates new Config from JSON-file if exists or from ENV
func NewTestConfig(fileName string, envPrefix string) (*Config, error) {
	var config *Config

	if f, err := os.Open(fileName); os.IsNotExist(err) {
		config, err = newConfigFromEnv(envPrefix)
		if err != nil {
			return nil, err
		}
	} else {
		if err != nil {
			return nil, err
		}
		defer f.Close()

		bytes, err := io.ReadAll(io.Reader(f))
		if err != nil {
			return nil, err
		}

		config, err = newConfigFromJSON(bytes)
		if err != nil {
			return nil, err
		}
	}

	if err := config.JwtConfig.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func newConfigFromJSON(bytes []byte) (*Config, error) {
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

func newConfigFromEnv(prefix string) (*Config, error) {
	pConfig := &Config{
		JwtConfig: *jwt.NewConfig("", ""),
		APIConfig: *api.NewConfiguration(),
	}

	pRoot := reflect.ValueOf(pConfig)
	root := pRoot.Elem()
	typeOfRoot := root.Type()

	for i := 0; i < typeOfRoot.NumField(); i++ {
		outer := root.Field(i)
		outerType := outer.Type()
		for j := 0; j < outerType.NumField(); j++ {
			inner := outer.Field(j)

			if inner.Kind() != reflect.String || !inner.CanSet() {
				continue
			}

			envName := fmt.Sprintf("%s_%s_%s", prefix,
				camelToUpperSnake(strings.Split(typeOfRoot.Field(i).Tag.Get("json"), ",")[0]),
				camelToUpperSnake(strings.Split(outerType.Field(j).Tag.Get("json"), ",")[0]))
			if newVal, exists := os.LookupEnv(envName); exists {
				fmt.Printf("Set '%s'\n", envName)
				inner.SetString(newVal)
			}
		}
	}

	return pConfig, nil
}

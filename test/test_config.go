package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"unicode"

	api "github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode"
	"github.com/aspose-barcode-cloud/aspose-barcode-cloud-go/barcode/jwt"
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

//Config is Configuration
type Config struct {
	JwtConfig jwt.Config        `json:"jwt"`
	APIConfig api.Configuration `json:"api"`
}

//NewConfig creates new Config from JSON-file
func NewConfig(fileName string) (*Config, error) {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return NewConfigFromEnv("TEST")
	}

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

//NewConfigFromEnv creates new Config from Environment
func NewConfigFromEnv(prefix string) (*Config, error) {
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

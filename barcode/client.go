package barcode

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	netUrl "net/url"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/oauth2"
)

const (
	PACKAGE_VERSION         = "4.2501.0"
	PACKAGE_NAME            = "go sdk"
	X_ASPOSE_CLIENT         = "x-aspose-client"
	X_ASPOSE_CLIENT_VERSION = "x-aspose-client-version"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

// APIClient manages communication with the Aspose.BarCode.Cloud v4.0 specification API v4.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	GenerateAPI *GenerateAPIService

	RecognizeAPI *RecognizeAPIService

	ScanAPI *ScanAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.GenerateAPI = (*GenerateAPIService)(&c.common)
	c.RecognizeAPI = (*RecognizeAPIService)(&c.common)
	c.ScanAPI = (*ScanAPIService)(&c.common)

	return c
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// ChangeBasePath to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams netUrl.Values,
	formParams netUrl.Values,
	fileName string,
	fileFieldName string,
	fileBytes []byte) (httpReq *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				// form value
				err = w.WriteField(k, iv)
				if err != nil {
					return nil, err
				}
			}
			w.Boundary()
		}

		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(fileFieldName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
		}

		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		err := w.Close()
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	url, err := netUrl.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		httpReq, err = http.NewRequest(method, url.String(), body)
	} else {
		httpReq, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		httpReq.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		httpReq.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	httpReq.Header.Add("User-Agent", c.cfg.UserAgent)
	httpReq.Header.Add(X_ASPOSE_CLIENT, PACKAGE_NAME)
	httpReq.Header.Add(X_ASPOSE_CLIENT_VERSION, PACKAGE_VERSION)

	if ctx != nil {
		// add context to the request
		httpReq = httpReq.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(httpReq)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			httpReq.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			httpReq.Header.Add("Authorization", "Bearer "+auth)
		}

		// JWT Authentication
		if tokSrc, ok := ctx.Value(ContextJWT).(oauth2.TokenSource); ok {
			jwtToken, err := tokSrc.Token()
			if err != nil {
				return nil, err
			}
			jwtToken.SetAuthHeader(httpReq)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		httpReq.Header.Add(header, value)
	}

	return httpReq, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.HasPrefix(contentType, "image/") {
		return nil
	}

	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}

	if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}

	return errors.New("undefined response type")
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// GenericAPIError Provides access to the body, error and model on returned errors.
type GenericAPIError struct {
	error      string
	text       string
	model      interface{}
	StatusCode int
}

// Error returns non-empty string if there was an error.
func (e GenericAPIError) Error() string {
	return e.error
}

// Text returns the text of the response
func (e GenericAPIError) Text() string {
	return e.text
}

// Model returns the unpacked model of the error
func (e GenericAPIError) Model() interface{} {
	return e.model
}

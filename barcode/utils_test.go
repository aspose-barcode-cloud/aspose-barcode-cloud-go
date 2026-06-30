package barcode

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type nullableValue[T comparable] interface {
	Get() *T
	Set(*T)
	IsSet() bool
	Unset()
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

func TestPointerHelpers(t *testing.T) {
	now := time.Date(2026, 6, 30, 12, 0, 0, 0, time.UTC)

	assert.Equal(t, true, *PtrBool(true))
	assert.Equal(t, 7, *PtrInt(7))
	assert.Equal(t, int32(8), *PtrInt32(8))
	assert.Equal(t, int64(9), *PtrInt64(9))
	assert.Equal(t, float32(1.25), *PtrFloat32(1.25))
	assert.Equal(t, 2.5, *PtrFloat64(2.5))
	assert.Equal(t, "value", *PtrString("value"))
	assert.Equal(t, now, *PtrTime(now))
}

func TestNullableHelpers(t *testing.T) {
	assertNullableValue(t, &NullableBool{}, true, "true")
	assertNullableValue(t, &NullableInt{}, 7, "7")
	assertNullableValue(t, &NullableInt32{}, int32(8), "8")
	assertNullableValue(t, &NullableInt64{}, int64(9), "9")
	assertNullableValue(t, &NullableFloat32{}, float32(1.25), "1.25")
	assertNullableValue(t, &NullableFloat64{}, 2.5, "2.5")
	assertNullableValue(t, &NullableString{}, "value", `"value"`)

	now := time.Date(2026, 6, 30, 12, 0, 0, 0, time.UTC)
	assertNullableValue(t, &NullableTime{}, now, `"2026-06-30T12:00:00Z"`)

	assert.Equal(t, false, *NewNullableBool(PtrBool(false)).Get())
	assert.Equal(t, 1, *NewNullableInt(PtrInt(1)).Get())
	assert.Equal(t, int32(2), *NewNullableInt32(PtrInt32(2)).Get())
	assert.Equal(t, int64(3), *NewNullableInt64(PtrInt64(3)).Get())
	assert.Equal(t, float32(4.5), *NewNullableFloat32(PtrFloat32(4.5)).Get())
	assert.Equal(t, 5.5, *NewNullableFloat64(PtrFloat64(5.5)).Get())
	assert.Equal(t, "created", *NewNullableString(PtrString("created")).Get())
	assert.Equal(t, now, *NewNullableTime(PtrTime(now)).Get())
}

func assertNullableValue[T comparable](t *testing.T, nullable nullableValue[T], value T, encoded string) {
	t.Helper()

	assert.False(t, nullable.IsSet())

	nullable.Set(&value)
	require.True(t, nullable.IsSet())
	require.NotNil(t, nullable.Get())
	assert.Equal(t, value, *nullable.Get())

	marshaled, err := nullable.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, encoded, string(marshaled))

	nullable.Unset()
	assert.False(t, nullable.IsSet())
	assert.Nil(t, nullable.Get())

	require.NoError(t, nullable.UnmarshalJSON([]byte(encoded)))
	require.True(t, nullable.IsSet())
	require.NotNil(t, nullable.Get())
	assert.Equal(t, value, *nullable.Get())
}

func TestNilAndStrictJSONHelpers(t *testing.T) {
	var ptr *int
	var values []string

	assert.True(t, IsNil(nil))
	assert.True(t, IsNil(ptr))
	assert.True(t, IsNil(values))
	assert.True(t, IsNil([2]int{}))
	assert.False(t, IsNil([2]int{1}))
	assert.False(t, IsNil("value"))

	var strictTarget struct {
		Name string `json:"name"`
	}
	decoder := newStrictDecoder([]byte(`{"name":"ok","extra":true}`))
	require.Error(t, decoder.Decode(&strictTarget))

	assert.EqualError(t, reportError("bad %s", "value"), "bad value")
}

func TestClientUtilityHelpers(t *testing.T) {
	assert.Equal(t, "", selectHeaderContentType(nil))
	assert.Equal(t, "text/plain", selectHeaderContentType([]string{"text/plain"}))
	assert.Equal(t, "", selectHeaderAccept(nil))
	assert.Equal(t, "image/png,text/plain", selectHeaderAccept([]string{"image/png", "text/plain"}))

	assert.Equal(t, "1|2|3", parameterToString([]int{1, 2, 3}, "pipes"))
	assert.Equal(t, "1 2 3", parameterToString([]int{1, 2, 3}, "ssv"))
	assert.Equal(t, "1\t2\t3", parameterToString([]int{1, 2, 3}, "tsv"))
	assert.Equal(t, "1,2,3", parameterToString([]int{1, 2, 3}, "csv"))
	assert.Equal(t, "value", parameterToString("value", ""))

	client := NewAPIClient(NewConfiguration())
	client.ChangeBasePath("https://changed.example")
	assert.Equal(t, "https://changed.example", client.cfg.BasePath)

	assert.Equal(t, "auth token", ContextOAuth2.String())

	apiError := GenericAPIError{error: "boom", text: "body"}
	assert.Equal(t, "boom", apiError.Error())
	assert.Equal(t, "body", apiError.Text())
}

func TestBodyAndDecodeHelpers(t *testing.T) {
	body, err := setBody(strings.NewReader("reader"), "text/plain")
	require.NoError(t, err)
	assert.Equal(t, "reader", body.String())

	body, err = setBody([]byte("bytes"), "application/octet-stream")
	require.NoError(t, err)
	assert.Equal(t, "bytes", body.String())

	body, err = setBody("string", "text/plain")
	require.NoError(t, err)
	assert.Equal(t, "string", body.String())

	value := "pointer"
	body, err = setBody(&value, "text/plain")
	require.NoError(t, err)
	assert.Equal(t, "pointer", body.String())

	body, err = setBody(map[string]string{"name": "value"}, "application/json")
	require.NoError(t, err)
	assert.JSONEq(t, `{"name":"value"}`, body.String())

	body, err = setBody(struct {
		XMLName xml.Name `xml:"root"`
		Name    string   `xml:"name"`
	}{Name: "value"}, "application/xml")
	require.NoError(t, err)
	assert.Contains(t, body.String(), "<name>value</name>")

	_, err = setBody(123, "application/octet-stream")
	require.EqualError(t, err, "invalid body type application/octet-stream")

	assert.Equal(t, "application/json; charset=utf-8", detectContentType(map[string]string{}))
	assert.Equal(t, "application/json; charset=utf-8", detectContentType([]string{"value"}))
	assert.Equal(t, "text/plain; charset=utf-8", detectContentType("value"))
	assert.Equal(t, "text/plain; charset=utf-8", detectContentType([]byte("value")))

	client := NewAPIClient(NewConfiguration())
	var image []byte
	require.NoError(t, client.decode(&image, []byte("image"), "image/png"))

	var jsonTarget map[string]string
	require.NoError(t, client.decode(&jsonTarget, []byte(`{"name":"value"}`), "application/json"))
	assert.Equal(t, "value", jsonTarget["name"])

	var xmlTarget struct {
		XMLName xml.Name `xml:"root"`
		Name    string   `xml:"name"`
	}
	require.NoError(t, client.decode(&xmlTarget, []byte(`<root><name>value</name></root>`), "application/xml"))
	assert.Equal(t, "value", xmlTarget.Name)

	require.EqualError(t, client.decode(&jsonTarget, []byte("value"), "text/plain"), "undefined response type")
}

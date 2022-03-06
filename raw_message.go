package payload

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// RawMessage extends basic json.RawMessage with extraction and manipulation methods of gjson/sjson.
type RawMessage json.RawMessage

// NewRawMessage creates new RawMessage using specified data mapping. Panics if underlying conversion failed.
// If you need json.RawMessage result use utility function NewJSON which do everything required under the hood.
// Provided data should contain flat mapping of path->value pairs where path is in dot syntax,
// such as "name.last" or "age". When the value is found it's returned immediately.
// A path is a series of keys separated by a dot. A key may contain special wildcard characters '*' and '?'.
// To access an array value use the index as the key.
// To get the number of elements in an array or to access a child path, use the '#' character.
// The dot and wildcard character can be escaped with '\'.
func NewRawMessage(data DataMap) (r *RawMessage) {
	var err error
	if r, err = MakeRawMessage(data); err != nil {
		panic(err)
	}

	return r
}

// MakeRawMessage creates new RawMessage using specified data mapping. Returns error if underlying conversion failed.
// If you need json.RawMessage with error instead use utility function MakeJSON.
func MakeRawMessage(data DataMap) (r *RawMessage, err error) {
	var jsonMessage json.RawMessage

	if jsonMessage, err = MakeJSON(data); err != nil {
		return nil, err
	}

	return (*RawMessage)(&jsonMessage), nil
}

// Wrap makes an RawMessage from provided json.RawMessage.
// It's a simple wrapper around RawMessage(message).
func Wrap(message json.RawMessage) (r *RawMessage) {
	rawMessage := RawMessage(message)
	return &rawMessage
}

// String returns string representation of RawData. Implements fmt.Stringer.
func (r RawMessage) String() string {
	return string(r)
}

// Get returns gjson.Result value taken from payload by specified path.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
func (r RawMessage) Get(key string) gjson.Result {
	return gjson.GetBytes(r, key)
}

// Set updates RawMessage adding or modifying specified key value.
func (r *RawMessage) Set(key string, value interface{}) (err error) {
	if *r, err = sjson.SetBytes(*r, key, value); err != nil {
		return fmt.Errorf("%v: set key: %w", Error, err)
	}

	return nil
}

// Update updates RawMessage in place adding or modifying keys from specified DataMap.
func (r *RawMessage) Update(data DataMap) error {
	updated, err := ExtendJSON(json.RawMessage(*r), data)
	if err != nil {
		return err
	}

	*r = RawMessage(updated)

	return nil
}

// Exists returns true if specified key path exists in request params.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
func (r RawMessage) Exists(path string) bool {
	return gjson.GetBytes(r, path).Exists()
}

// GetString returns string value taken from request params by specified path.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
// If specified key empty or not exists returns empty string.
// To check if key exists use Exists or get raw result with Get.
func (r RawMessage) GetString(path string) string {
	return gjson.GetBytes(r, path).String()
}

// GetInt returns int value taken from request params by specified path.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
// If specified key empty or not exists returns 0.
// To check if key exists use Exists or get raw result with Get.
func (r RawMessage) GetInt(path string) int {
	return int(gjson.GetBytes(r, path).Int())
}

// GetFloat64 returns float64 value taken from request params by specified path.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
// If specified key empty or not exists returns 0.
// To check if key exists use Exists or get raw result with Get.
func (r RawMessage) GetFloat64(path string) float64 {
	return gjson.GetBytes(r, path).Float()
}

// GetBool returns bool value taken from request params by specified path.
// Path may be search query, see https://github.com/tidwall/gjson for detailed info.
// If specified key empty or not exists returns false.
// To check if key exists use Exists or get raw result with Get.
// NOTE: for non-boolean values if called for numeric field it returns false if value=0 and true otherwise.
// For string field it returns false if string is empty, equals "0" or "false" and true otherwise.
func (r RawMessage) GetBool(path string) bool {
	return gjson.GetBytes(r, path).Bool()
}

// MarshalJSON does json marshalling of raw message. Implements json.Marshaler.
func (r RawMessage) MarshalJSON() ([]byte, error) {
	return r, nil
}

// UnmarshalJSON does json unmarshalling of raw message. Implements json.Unmarshaler.
func (r *RawMessage) UnmarshalJSON(data []byte) error {
	*r = data

	return nil
}

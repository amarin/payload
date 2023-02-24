package payload

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/sjson"
)

// ExtendJSON extend existing raw JSON message of json.RawMessage with specified data,
// doing keys unpacking by dot separator, f.e. pair "a.b.c": true became a>b>c=true.
// If creating new message from keys use normalizeKeys.
func ExtendJSON(origJSON json.RawMessage, data DataMap) (result json.RawMessage, err error) {
	if data == nil {
		return origJSON, nil
	}
	result = origJSON
	for k, v := range data {
		if result, err = sjson.SetBytes(result, k, v); err != nil {
			return nil, fmt.Errorf("%w: set %v: %v", Error, k, err)
		}
	}

	return result, nil
}

// MustExtendJSON extend existing raw JSON message of json.RawMessage with specified data,
// doing keys unpacking by dot separator, f.e. pair "a.b.c": true became a>b>c=true.
// If creating new message from keys use normalizeKeys.
// Panics if error happened during raw JSON updating.
func MustExtendJSON(origJSON json.RawMessage, data DataMap) (result json.RawMessage) {
	var err error

	if result, err = ExtendJSON(origJSON, data); err != nil {
		panic(err)
	}

	return result
}

// MakeJSON prepares json.RawMessage suitable to use as request data or response result.
// If any key contains dot it became multilayer key f.e. pair "a.b.c": true became a>b>c=true.
// It is a shorthand to call ExtendJSON([]byte(EmptyJSONObject), ...).
func MakeJSON(data DataMap) (result json.RawMessage, err error) {
	return ExtendJSON([]byte(EmptyJSONObject), data)
}

// NewJSON prepares json.RawMessage suitable to use as request data or response result.
// If any key contains dot it became multilayer key f.e. pair "a.b.c": true became a>b>c=true.
// It is a shorthand to call ExtendJSON([]byte(EmptyJSONObject), ...). Panics if error happened during raw JSON creation.
func NewJSON(data DataMap) json.RawMessage {
	return MustExtendJSON([]byte(EmptyJSONObject), data)
}

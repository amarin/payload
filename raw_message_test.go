package payload_test

import (
	"fmt"
	"github.com/amarin/payload"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeRawMessage(t *testing.T) {
	for _, tt := range []struct {
		name    string
		data    payload.DataMap
		wantR   string
		wantErr bool
	}{
		{
			"simple_one",
			payload.DataMap{"one": 1},
			`{"one":1}`,
			false,
		},
		{
			"str_list_and_str",
			payload.DataMap{"subsystems": []string{"aaa", "bbb"}, "apiKey": "TEST_API_KEY"},
			`{"apiKey":"TEST_API_KEY","subsystems":["aaa", "bbb"]}`,
			false,
		},
		{
			"str_list_and_obj",
			payload.DataMap{
				"list": []string{"a1", "a2"},
				"obj": map[string]interface{}{
					"attr1": 1,
					"attr2": true,
					"attr3": "str",
					"attr4": 1.2,
				},
			},
			`{"list":["a1", "a2"], "obj":{"attr1": 1,"attr2": true, "attr3": "str","attr4": 1.2}}`,
			false,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt := tt
			gotR, err := payload.MakeRawMessage(tt.data)
			require.Equalf(t, tt.wantErr, err != nil, "MakeRawMessage() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotNil(t, gotR)
			require.JSONEq(t, tt.wantR, string(*gotR))
		})
	}
}

func TestNewRawMessage(t *testing.T) {
	tests := []struct {
		name  string
		args  payload.DataMap
		wantR string
	}{
		{"make_empty_raw_message", nil, payload.EmptyJSONObject},
		{"make_flat_raw_message", payload.DataMap{"a": "b", "c": 1}, `{"a":"b","c":1}`},
		{"make_tree_raw_message", payload.DataMap{"a": payload.DataMap{"c": 1}}, `{"a":{"c":1}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.JSONEq(t, string(tt.wantR), string(*payload.NewRawMessage(tt.args)))
		})
	}
}

func ExampleNewRawMessage() {
	// make new empty RawMessage
	rawJSON := payload.NewRawMessage(nil)
	fmt.Println(rawJSON.String())

	// make new RawMessage with single string value
	newRawJSON := payload.NewRawMessage(payload.DataMap{"stringKey": "stringValue"})
	fmt.Println(newRawJSON.String())

	// take string value using known key
	stringKey := newRawJSON.GetString("stringKey")
	fmt.Println(stringKey)

	// Output:
	// {}
	// {"stringKey":"stringValue"}
	// stringValue
}

func ExampleMakeRawMessage() {
	// make new empty MakeMessage with error handling
	rawJSON, err := payload.MakeRawMessage(nil)
	if err != nil {
		log.Fatal("MakeRawMessage:", err)
	}
	fmt.Println(rawJSON.String())

	// make new RawMessage with single string value
	newRawJSON, err := payload.MakeRawMessage(payload.DataMap{"intKey": 10})
	if err != nil {
		log.Fatal("MakeRawMessage:", err)
	}
	fmt.Println(newRawJSON.String())

	// take string value using known key
	stringKey := newRawJSON.GetInt("intKey")
	fmt.Println(stringKey)

	// Output:
	// {}
	// {"intKey":10}
	// 10
}

func TestWrap(t *testing.T) {
	gotR := payload.Wrap([]byte("{}"))
	require.IsType(t, &payload.RawMessage{}, gotR)
}

func ExampleWrap() {
	rawMessage := payload.Wrap([]byte(`{"intKey":1,"boolKey":true,"stringKey":"iAmAString"}`))
	fmt.Println(
		rawMessage.GetInt("intKey"),
		rawMessage.GetBool("boolKey"),
		rawMessage.GetString("stringKey"),
	)

	// Output
	// 1 true iAmAString
}

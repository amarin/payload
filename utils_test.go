package payload_test

import (
	"encoding/json"
	"testing"

	"git.media-tel.ru/railgo/payload"
	"github.com/stretchr/testify/require"
)

func TestExtendJSON(t *testing.T) {
	for _, tt := range []struct {
		name       string
		origJSON   string
		data       payload.DataMap
		wantResult string
		wantErr    bool
	}{
		{
			"add_single_key",
			payload.EmptyJSONObject,
			payload.DataMap{"one": 1},
			`{"one":1}`,
			false},
		{
			"add_couple_of_key",
			payload.EmptyJSONObject,
			payload.DataMap{"one": 1, "two": "два"},
			`{"two":"два","one":1}`,
			false},
		{
			"add_empty_map_changes_nothing",
			`{"one":1, "two":"два"}`,
			payload.DataMap{},
			`{"two":"два","one":1}`,
			false},
		{
			"add_nil_map_is_not_error",
			`{"one":1, "two":"два"}`,
			nil,
			`{"two":"два","one":1}`,
			false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := payload.ExtendJSON(json.RawMessage(tt.origJSON), tt.data)
			require.Equal(t, err != nil, tt.wantErr)
			if err != nil {
				return
			}

			require.JSONEq(t, tt.wantResult, string(gotResult))
		})
	}
}

package payload

import "errors"

// EmptyJSONObject contains empty JSON object string representation.
const EmptyJSONObject = "{}"

// Error indicates payload routines errors.
var Error = errors.New("encode")

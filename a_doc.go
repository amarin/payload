// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package payload provides utility type and functions to deal with JSON data stored as byte array.
It wraps and uses awesome github.com/tidwall/sjson and github.com/tidwall/sjson packages and its types.

The main goal is to provide better switch from native json.RawMessage into gjson methods
and from sjson to native json.RawMessage.

To solve this the RawMessage type introduced.
It simply wraps json.RawMessage but provides proxy methods to gjson functions to get expected keys data
as well as proxy methods to sjson functions to modify underlying message if required.

To simplify operations in case many keys should be set simultaneously a DataMap type is provided.
It can be used with NewRawMessage or MakeRawMessage constructors as well as RawMessage.Update argument.
It wraps map[string]interface{} and provides DataMap.Update methods.

IMPORTANT: see NewRawMessage description to understand difference of key and path.

Use RawMessage

To build raw message from scratch use NewRawMessage or MakeRawMessage constructor
or simply wrap existing json.RawMessage instance with Wrap method.

See github.com/tidwall/gjson/ documentation to details of path usage in Get and Set functions.
*/
package payload

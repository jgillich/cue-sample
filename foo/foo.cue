package foo

import "encoding/json"

#foo: string @tag(foo)
json.Unmarshal(#foo)

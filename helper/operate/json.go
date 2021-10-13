package operateHelper

import "encoding/json"

func JSONEncodeToString(v interface{}) string {
    return string(JSONEncodeToByte(v))
}

func JSONEncodeToByte(v interface{}) (b []byte) {
    b, _ = json.Marshal(v)
    return
}

package output


import (
"encoding/json"
"os"
)


func PrintJSON(v any) {
enc := json.NewEncoder(os.Stdout)
enc.SetIndent("", " ")
enc.Encode(v)
}

package debug_dump_json

import "os"
import "encoding/json"

func Dump(any interface{}) {
  b, err := json.Marshal(any)
  if err != nil { panic(err) }
  
  var buf bytes.Buffer;
  err = json.Indent(&buf, b, "", "  ")
  if err != nil { panic(err) }
  
  _, _ = buf.WriteTo(os.Stderr)
  _, _ = os.Stderr.WriteString("\n")
}
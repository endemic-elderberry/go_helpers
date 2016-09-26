package exit_same_code

import "os"
import "regexp"
import "strconv"

var match_ecode *regexp.Regexp;

func init(){
  match_ecode = regexp.MustCompile(`^exit (?:code|status) (\d+)$`)
}

func GetCode(err error) int {
  if err==nil {
    return 0
  }
  match := match_ecode.FindStringSubmatch(err.Error())

  if match == nil { return 0 }

  i, _ := strconv.Atoi( match[1] )

  return i
}
func ExitSameCode(err error){
  code := GetCode(err)
  os.Exit( code )
}
func ExitSameCodeIfFailure(err error){
  if err==nil { return }

  ExitSameCode(err)
}

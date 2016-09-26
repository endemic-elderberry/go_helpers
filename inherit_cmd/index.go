package inherit_cmd

import "os/exec"
import "os"

func Inherit(cmd *exec.Cmd){
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
}

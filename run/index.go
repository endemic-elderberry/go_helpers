package run

import "os"
import "os/exec"
import "io"

var writeTo io.Writer;
var writeToErr io.Writer
var readFrom io.Reader;

func RunInherit(executable string, args... string) error {
  c := exec.Command(executable,args...)

  if writeTo == nil {
    c.Stdout = os.Stdout;
  } else {
    c.Stdout = writeTo
  }
  if writeToErr == nil {
    c.Stderr = os.Stderr;
  } else {
    c.Stderr = writeToErr
  }
  if readFrom == nil {
    c.Stdin = os.Stdin;
  } else {
    c.Stdin = readFrom
  }
  return c.Run()
}

func MustRunInherit(executable string, args... string) {
  err := RunInherit(executable, args...)
  if err != nil {
    panic(err)
  }
}


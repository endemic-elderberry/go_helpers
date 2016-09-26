package exit_same_code

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "os/exec"
)


func TestMatcher(t *testing.T) {
  err := exec.Command("false").Run()

  assert.Equal(t, GetCode(err), 1, "false should give 1 on unix systems")
}


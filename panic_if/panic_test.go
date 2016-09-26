package panic_if


import "testing"
import "errors"
import "github.com/stretchr/testify/assert"

func TestPanic(t *testing.T) {
  assert := assert.New(t)
  
  assert.Panics(func () {
    PanicIf(errors.New("Something Went Wrong"))
  }, "Should panic if an error is given")
  assert.NotPanics(func () {
    PanicIf(nil)
  }, "Should not panic if an error is nil")
}
package run

import "testing"
import "github.com/stretchr/testify/assert"
// import "os"

type writeStore struct{
  Data string
}
func (store *writeStore) Write(b []byte) (int, error) {
  store.Data += string(b)
  
  return len(b),nil
}

func TestRun(t *testing.T) {
  assert := assert.New(t)
  
  err := RunInherit("nonexistantexecutable")
  assert.NotNil(err)
  
  
  myStore := &writeStore{""}
  writeTo = myStore;
  
  err = RunInherit("echo", "hi")
  assert.Nil(err)
  
  assert.Equal("hi\n", myStore.Data, "echo output captured")
  
  assert.Panics(func () {
    MustRunInherit("nonexistantexecutable")
  },"Should panic if ran with Must~")
  assert.NotPanics(func () {
    MustRunInherit("echo", "hi")
  },"Should not panic if valid")
  
  writeTo = nil;
}


package keys

import "testing"
import "math/rand"
import "github.com/stretchr/testify/assert"

func TestKeyGet(t *testing.T) {
  assert := assert.New(t)
  
  testMap := make(map[interface{}]interface{})
  originalMap := make(map[interface{}]interface{})
  
  for i := 0 ; i < 100 ; i++ {
    key := rand.Int31()
    testMap[key] = true
    originalMap[key] = true
  }
  
  notInMap := make(map[interface{}] interface{})
  
  i := 0
  for k, _ := range originalMap {
    if i % 2 == 0 {
      delete(testMap, k)
      notInMap[k] = true
    }
    i++
  }
  
  keyList := Keys(testMap)
  
  inKeyList := make(map[interface{}]interface{})
  
  for _, k := range keyList {
    inKeyList[k] = true
  }
  
  for k, _ := range originalMap  {
    _, okTest := testMap[k]
    _, okKeyList:= inKeyList[k]
    _, okNot := notInMap[k]
    
    if okNot {
      assert.True(!okTest, "Was deleted")
      assert.True(!okKeyList, "Not in output list")
    } else {
      assert.True(okTest, "Was in testMap")
      assert.True(okKeyList, "In output keyList")
    }
  }
}

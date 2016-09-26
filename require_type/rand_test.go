package require_type

import "testing"
import "math/rand"
import "github.com/stretchr/testify/assert"

//go:generate node generate/generateTest.js


func generateFromExpFloat64 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.ExpFloat64()
  }
  return interfaceList
}
func TestBadCastArrayFromExpFloat64(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromExpFloat64()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_float64(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromExpFloat64(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromExpFloat64()
  asTypedArray := CastToArrayOf_float64(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromFloat32 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Float32()
  }
  return interfaceList
}
func TestBadCastArrayFromFloat32(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromFloat32()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_float32(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromFloat32(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromFloat32()
  asTypedArray := CastToArrayOf_float32(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromFloat64 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Float64()
  }
  return interfaceList
}
func TestBadCastArrayFromFloat64(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromFloat64()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_float64(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromFloat64(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromFloat64()
  asTypedArray := CastToArrayOf_float64(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromInt () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Int()
  }
  return interfaceList
}
func TestBadCastArrayFromInt(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromInt()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_int(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromInt(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromInt()
  asTypedArray := CastToArrayOf_int(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromInt31 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Int31()
  }
  return interfaceList
}
func TestBadCastArrayFromInt31(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromInt31()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_int32(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromInt31(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromInt31()
  asTypedArray := CastToArrayOf_int32(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromInt63 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Int63()
  }
  return interfaceList
}
func TestBadCastArrayFromInt63(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromInt63()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_int64(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromInt63(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromInt63()
  asTypedArray := CastToArrayOf_int64(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromNormFloat64 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.NormFloat64()
  }
  return interfaceList
}
func TestBadCastArrayFromNormFloat64(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromNormFloat64()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_float64(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromNormFloat64(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromNormFloat64()
  asTypedArray := CastToArrayOf_float64(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}



func generateFromUint32 () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.Uint32()
  }
  return interfaceList
}
func TestBadCastArrayFromUint32(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromUint32()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_uint32(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromUint32(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromUint32()
  asTypedArray := CastToArrayOf_uint32(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}


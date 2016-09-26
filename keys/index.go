package keys

import "reflect"

func Keys(m map[interface{}]interface{}) []interface{} {
  mValue := reflect.ValueOf(m)
  
  mKind := mValue.Kind()
  if mKind != reflect.Map {
    panic("Not a map")
  }
  
  mKeys := mValue.MapKeys()
  
  iSlice := make([]interface{}, len(mKeys))
  for i := 0; i < len(mKeys); i++ {
    vCurrent := mKeys[ i ].Interface()
    iSlice[ i ] = vCurrent
  }
  return iSlice
}

type KeyGetter interface {
  Keys(map[interface{}]interface{}) []interface{}  
}
type KeyGetterStruct struct{};

func (kgs KeyGetterStruct) Keys(m map[interface{}]interface{}) ([]interface{}) {
  return Keys(m)
}
func MakeKeyGetter() KeyGetter {
  return new(KeyGetterStruct)
}


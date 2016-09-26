package require_type

import "fmt"
//go:generate node generate/generate.js

type MapFetcher struct{
  SourceMap map[interface{}]interface{}
}

func MakeMapFetcher(m map[interface{}]interface{}) MapFetcher {
  return MapFetcher{
    SourceMap: m,
  };
}

func (m MapFetcher) FetchAny(key interface{}) interface{} {
  var v_temp interface{};
  v_temp, _ = m.SourceMap[key]
  return v_temp;
}
func (m MapFetcher) RequireAny(key interface{}) interface{} {
  var v_temp interface{};
  var vok bool;
  v_temp, vok = m.SourceMap[key]
  if !vok {
    panic(fmt.Sprintf("%s not found", key))
  }
  return v_temp;
}


func (m MapFetcher) Fetch_bool(key interface{}) bool {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(bool);
  return v_as
}
func (m MapFetcher) Require_bool(key interface{}) bool {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(bool);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_bool(l []interface{}) []bool {
  a_to := make([]bool,len(l));
  var v_as bool;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(bool);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "bool"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_bool(key interface{}, m map[interface{}]interface{}) bool {
  return MakeMapFetcher(m).Fetch_bool(key);
}
func RequireFromMap_bool(key interface{}, m map[interface{}]interface{}) bool {
  return MakeMapFetcher(m).Require_bool(key);
}


func (m MapFetcher) Fetch_byte(key interface{}) byte {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(byte);
  return v_as
}
func (m MapFetcher) Require_byte(key interface{}) byte {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(byte);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_byte(l []interface{}) []byte {
  a_to := make([]byte,len(l));
  var v_as byte;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(byte);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "byte"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_byte(key interface{}, m map[interface{}]interface{}) byte {
  return MakeMapFetcher(m).Fetch_byte(key);
}
func RequireFromMap_byte(key interface{}, m map[interface{}]interface{}) byte {
  return MakeMapFetcher(m).Require_byte(key);
}


func (m MapFetcher) Fetch_complex128(key interface{}) complex128 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(complex128);
  return v_as
}
func (m MapFetcher) Require_complex128(key interface{}) complex128 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(complex128);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_complex128(l []interface{}) []complex128 {
  a_to := make([]complex128,len(l));
  var v_as complex128;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(complex128);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "complex128"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_complex128(key interface{}, m map[interface{}]interface{}) complex128 {
  return MakeMapFetcher(m).Fetch_complex128(key);
}
func RequireFromMap_complex128(key interface{}, m map[interface{}]interface{}) complex128 {
  return MakeMapFetcher(m).Require_complex128(key);
}


func (m MapFetcher) Fetch_complex64(key interface{}) complex64 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(complex64);
  return v_as
}
func (m MapFetcher) Require_complex64(key interface{}) complex64 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(complex64);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_complex64(l []interface{}) []complex64 {
  a_to := make([]complex64,len(l));
  var v_as complex64;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(complex64);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "complex64"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_complex64(key interface{}, m map[interface{}]interface{}) complex64 {
  return MakeMapFetcher(m).Fetch_complex64(key);
}
func RequireFromMap_complex64(key interface{}, m map[interface{}]interface{}) complex64 {
  return MakeMapFetcher(m).Require_complex64(key);
}


func (m MapFetcher) Fetch_error(key interface{}) error {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(error);
  return v_as
}
func (m MapFetcher) Require_error(key interface{}) error {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(error);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_error(l []interface{}) []error {
  a_to := make([]error,len(l));
  var v_as error;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(error);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "error"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_error(key interface{}, m map[interface{}]interface{}) error {
  return MakeMapFetcher(m).Fetch_error(key);
}
func RequireFromMap_error(key interface{}, m map[interface{}]interface{}) error {
  return MakeMapFetcher(m).Require_error(key);
}


func (m MapFetcher) Fetch_float32(key interface{}) float32 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(float32);
  return v_as
}
func (m MapFetcher) Require_float32(key interface{}) float32 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(float32);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_float32(l []interface{}) []float32 {
  a_to := make([]float32,len(l));
  var v_as float32;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(float32);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "float32"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_float32(key interface{}, m map[interface{}]interface{}) float32 {
  return MakeMapFetcher(m).Fetch_float32(key);
}
func RequireFromMap_float32(key interface{}, m map[interface{}]interface{}) float32 {
  return MakeMapFetcher(m).Require_float32(key);
}


func (m MapFetcher) Fetch_float64(key interface{}) float64 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(float64);
  return v_as
}
func (m MapFetcher) Require_float64(key interface{}) float64 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(float64);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_float64(l []interface{}) []float64 {
  a_to := make([]float64,len(l));
  var v_as float64;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(float64);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "float64"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_float64(key interface{}, m map[interface{}]interface{}) float64 {
  return MakeMapFetcher(m).Fetch_float64(key);
}
func RequireFromMap_float64(key interface{}, m map[interface{}]interface{}) float64 {
  return MakeMapFetcher(m).Require_float64(key);
}


func (m MapFetcher) Fetch_int(key interface{}) int {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(int);
  return v_as
}
func (m MapFetcher) Require_int(key interface{}) int {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(int);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_int(l []interface{}) []int {
  a_to := make([]int,len(l));
  var v_as int;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(int);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "int"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_int(key interface{}, m map[interface{}]interface{}) int {
  return MakeMapFetcher(m).Fetch_int(key);
}
func RequireFromMap_int(key interface{}, m map[interface{}]interface{}) int {
  return MakeMapFetcher(m).Require_int(key);
}


func (m MapFetcher) Fetch_int16(key interface{}) int16 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(int16);
  return v_as
}
func (m MapFetcher) Require_int16(key interface{}) int16 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(int16);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_int16(l []interface{}) []int16 {
  a_to := make([]int16,len(l));
  var v_as int16;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(int16);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "int16"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_int16(key interface{}, m map[interface{}]interface{}) int16 {
  return MakeMapFetcher(m).Fetch_int16(key);
}
func RequireFromMap_int16(key interface{}, m map[interface{}]interface{}) int16 {
  return MakeMapFetcher(m).Require_int16(key);
}


func (m MapFetcher) Fetch_int32(key interface{}) int32 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(int32);
  return v_as
}
func (m MapFetcher) Require_int32(key interface{}) int32 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(int32);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_int32(l []interface{}) []int32 {
  a_to := make([]int32,len(l));
  var v_as int32;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(int32);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "int32"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_int32(key interface{}, m map[interface{}]interface{}) int32 {
  return MakeMapFetcher(m).Fetch_int32(key);
}
func RequireFromMap_int32(key interface{}, m map[interface{}]interface{}) int32 {
  return MakeMapFetcher(m).Require_int32(key);
}


func (m MapFetcher) Fetch_int64(key interface{}) int64 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(int64);
  return v_as
}
func (m MapFetcher) Require_int64(key interface{}) int64 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(int64);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_int64(l []interface{}) []int64 {
  a_to := make([]int64,len(l));
  var v_as int64;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(int64);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "int64"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_int64(key interface{}, m map[interface{}]interface{}) int64 {
  return MakeMapFetcher(m).Fetch_int64(key);
}
func RequireFromMap_int64(key interface{}, m map[interface{}]interface{}) int64 {
  return MakeMapFetcher(m).Require_int64(key);
}


func (m MapFetcher) Fetch_int8(key interface{}) int8 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(int8);
  return v_as
}
func (m MapFetcher) Require_int8(key interface{}) int8 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(int8);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_int8(l []interface{}) []int8 {
  a_to := make([]int8,len(l));
  var v_as int8;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(int8);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "int8"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_int8(key interface{}, m map[interface{}]interface{}) int8 {
  return MakeMapFetcher(m).Fetch_int8(key);
}
func RequireFromMap_int8(key interface{}, m map[interface{}]interface{}) int8 {
  return MakeMapFetcher(m).Require_int8(key);
}


func (m MapFetcher) Fetch_rune(key interface{}) rune {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(rune);
  return v_as
}
func (m MapFetcher) Require_rune(key interface{}) rune {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(rune);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_rune(l []interface{}) []rune {
  a_to := make([]rune,len(l));
  var v_as rune;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(rune);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "rune"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_rune(key interface{}, m map[interface{}]interface{}) rune {
  return MakeMapFetcher(m).Fetch_rune(key);
}
func RequireFromMap_rune(key interface{}, m map[interface{}]interface{}) rune {
  return MakeMapFetcher(m).Require_rune(key);
}


func (m MapFetcher) Fetch_string(key interface{}) string {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(string);
  return v_as
}
func (m MapFetcher) Require_string(key interface{}) string {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(string);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_string(l []interface{}) []string {
  a_to := make([]string,len(l));
  var v_as string;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(string);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "string"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_string(key interface{}, m map[interface{}]interface{}) string {
  return MakeMapFetcher(m).Fetch_string(key);
}
func RequireFromMap_string(key interface{}, m map[interface{}]interface{}) string {
  return MakeMapFetcher(m).Require_string(key);
}


func (m MapFetcher) Fetch_uint(key interface{}) uint {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uint);
  return v_as
}
func (m MapFetcher) Require_uint(key interface{}) uint {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uint);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uint(l []interface{}) []uint {
  a_to := make([]uint,len(l));
  var v_as uint;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uint);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uint"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uint(key interface{}, m map[interface{}]interface{}) uint {
  return MakeMapFetcher(m).Fetch_uint(key);
}
func RequireFromMap_uint(key interface{}, m map[interface{}]interface{}) uint {
  return MakeMapFetcher(m).Require_uint(key);
}


func (m MapFetcher) Fetch_uint16(key interface{}) uint16 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uint16);
  return v_as
}
func (m MapFetcher) Require_uint16(key interface{}) uint16 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uint16);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uint16(l []interface{}) []uint16 {
  a_to := make([]uint16,len(l));
  var v_as uint16;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uint16);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uint16"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uint16(key interface{}, m map[interface{}]interface{}) uint16 {
  return MakeMapFetcher(m).Fetch_uint16(key);
}
func RequireFromMap_uint16(key interface{}, m map[interface{}]interface{}) uint16 {
  return MakeMapFetcher(m).Require_uint16(key);
}


func (m MapFetcher) Fetch_uint32(key interface{}) uint32 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uint32);
  return v_as
}
func (m MapFetcher) Require_uint32(key interface{}) uint32 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uint32);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uint32(l []interface{}) []uint32 {
  a_to := make([]uint32,len(l));
  var v_as uint32;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uint32);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uint32"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uint32(key interface{}, m map[interface{}]interface{}) uint32 {
  return MakeMapFetcher(m).Fetch_uint32(key);
}
func RequireFromMap_uint32(key interface{}, m map[interface{}]interface{}) uint32 {
  return MakeMapFetcher(m).Require_uint32(key);
}


func (m MapFetcher) Fetch_uint64(key interface{}) uint64 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uint64);
  return v_as
}
func (m MapFetcher) Require_uint64(key interface{}) uint64 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uint64);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uint64(l []interface{}) []uint64 {
  a_to := make([]uint64,len(l));
  var v_as uint64;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uint64);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uint64"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uint64(key interface{}, m map[interface{}]interface{}) uint64 {
  return MakeMapFetcher(m).Fetch_uint64(key);
}
func RequireFromMap_uint64(key interface{}, m map[interface{}]interface{}) uint64 {
  return MakeMapFetcher(m).Require_uint64(key);
}


func (m MapFetcher) Fetch_uint8(key interface{}) uint8 {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uint8);
  return v_as
}
func (m MapFetcher) Require_uint8(key interface{}) uint8 {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uint8);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uint8(l []interface{}) []uint8 {
  a_to := make([]uint8,len(l));
  var v_as uint8;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uint8);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uint8"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uint8(key interface{}, m map[interface{}]interface{}) uint8 {
  return MakeMapFetcher(m).Fetch_uint8(key);
}
func RequireFromMap_uint8(key interface{}, m map[interface{}]interface{}) uint8 {
  return MakeMapFetcher(m).Require_uint8(key);
}


func (m MapFetcher) Fetch_uintptr(key interface{}) uintptr {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(uintptr);
  return v_as
}
func (m MapFetcher) Require_uintptr(key interface{}) uintptr {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(uintptr);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_uintptr(l []interface{}) []uintptr {
  a_to := make([]uintptr,len(l));
  var v_as uintptr;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(uintptr);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "uintptr"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_uintptr(key interface{}, m map[interface{}]interface{}) uintptr {
  return MakeMapFetcher(m).Fetch_uintptr(key);
}
func RequireFromMap_uintptr(key interface{}, m map[interface{}]interface{}) uintptr {
  return MakeMapFetcher(m).Require_uintptr(key);
}


func (m MapFetcher) Fetch_array_interface(key interface{}) []interface{} {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.([]interface{});
  return v_as
}
func (m MapFetcher) Require_array_interface(key interface{}) []interface{} {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.([]interface{});
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_array_interface(l []interface{}) [][]interface{} {
  a_to := make([][]interface{},len(l));
  var v_as []interface{};
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.([]interface{});
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "[]interface{}"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_array_interface(key interface{}, m map[interface{}]interface{}) []interface{} {
  return MakeMapFetcher(m).Fetch_array_interface(key);
}
func RequireFromMap_array_interface(key interface{}, m map[interface{}]interface{}) []interface{} {
  return MakeMapFetcher(m).Require_array_interface(key);
}


func (m MapFetcher) Fetch_map_interface_interface(key interface{}) map[interface{}]interface{} {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(map[interface{}]interface{});
  return v_as
}
func (m MapFetcher) Require_map_interface_interface(key interface{}) map[interface{}]interface{} {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(map[interface{}]interface{});
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_map_interface_interface(l []interface{}) []map[interface{}]interface{} {
  a_to := make([]map[interface{}]interface{},len(l));
  var v_as map[interface{}]interface{};
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(map[interface{}]interface{});
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "map[interface{}]interface{}"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_map_interface_interface(key interface{}, m map[interface{}]interface{}) map[interface{}]interface{} {
  return MakeMapFetcher(m).Fetch_map_interface_interface(key);
}
func RequireFromMap_map_interface_interface(key interface{}, m map[interface{}]interface{}) map[interface{}]interface{} {
  return MakeMapFetcher(m).Require_map_interface_interface(key);
}

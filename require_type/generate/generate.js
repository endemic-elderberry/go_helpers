
var fs = require ("fs");

process.chdir(__dirname);

var fw = fs.createWriteStream("../assert_to.go")

//https://golang.org/pkg/builtin/

var types = require('./getTypeInfo.js');

types.push('[]interface{}', 'map[interface{}]interface{}');


fw.write(`package require_type

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
`);

types.forEach(function (typeto) {
  var typefname = typeto.replace(/\{\}/g,"");
  typefname = typefname.replace(/\[\]/g,"array_");
  typefname = typefname.replace(/[\[\]]/g,"_");
  
  fw.write(`

func (m MapFetcher) Fetch_` + typefname + `(key interface{}) ` + typeto + ` {
  v_temp := m.FetchAny(key)
  v_as, _ := v_temp.(`+typeto+`);
  return v_as
}
func (m MapFetcher) Require_` + typefname + `(key interface{}) ` + typeto + ` {
  v_temp := m.RequireAny(key)
  v_as, vok := v_temp.(`+typeto+`);
  if !vok {
    panic(fmt.Sprintf("%s not correct type, is %T", key, key) )
  }
  return v_as
}
func CastToArrayOf_`+typefname+`(l []interface{}) []`+typeto+` {
  a_to := make([]`+typeto+`,len(l));
  var v_as `+typeto+`;
  var v_temp interface{};
  var vok bool;
  
  for i:=0 ; i < len(l) ; i++ {
    v_temp = l[i]
    
    v_as, vok = v_temp.(`+typeto+`);
    if !vok {
      panic(fmt.Sprintf("could not cast item at index %d to %s", i, "`+typeto+`"))
    }
    a_to[i] = v_as
  }
  return a_to
}
func FetchFromMap_`+typefname+`(key interface{}, m map[interface{}]interface{}) `+typeto+` {
  return MakeMapFetcher(m).Fetch_` + typefname + `(key);
}
func RequireFromMap_`+typefname+`(key interface{}, m map[interface{}]interface{}) `+typeto+` {
  return MakeMapFetcher(m).Require_` + typefname + `(key);
}
`);

});

fw.end();



var fs = require ("fs");
var child_process = require('child_process');

process.chdir(__dirname);

var types = require('./getRandInfo.js');
var fw = fs.createWriteStream("../rand_test.go")

fw.write(`package require_type

import "testing"
import "math/rand"
import "github.com/stretchr/testify/assert"

//go:generate node generate/generateTest.js
`);

var template = `

func generateFromFROMFUNC () []interface{} {
  interfaceList := make([]interface{}, 10)
  
  for i := 0 ; i < len(interfaceList) ; i++ {
    interfaceList[i] = rand.FROMFUNC()
  }
  return interfaceList
}
func TestBadCastArrayFromFROMFUNC(t *testing.T) {
  assert := assert.New(t)
  
  badInterfaceList := generateFromFROMFUNC()
  badInterfaceList = append(badInterfaceList, "Yep", 1)
  
  assert.Panics(func () {
    _ = CastToArrayOf_WHICHTYPE(badInterfaceList)
  }, "Should panic on bad input")
}
func TestWorkingCastArrayFromFROMFUNC(t *testing.T) {
  assert := assert.New(t)
  
  interfaceList := generateFromFROMFUNC()
  asTypedArray := CastToArrayOf_WHICHTYPE(interfaceList)
  
  for i := 0 ; i < len(asTypedArray) ; i++ {
    assert.Equal(asTypedArray[i], interfaceList[i], "Values should be identical")
  }
}

`;

types.forEach(function (type) {
  var replaced = template.replace(/FROMFUNC/g, type.generateFunc).replace(/WHICHTYPE/g, type.primitiveType);
  
  fw.write(replaced);


});

fw.end();

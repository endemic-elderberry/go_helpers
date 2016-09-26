
var fs = require('fs');
var path = require('path');
var child_process = require('child_process');
var util = require('util');


process.chdir(__dirname);

var files = fs.readdirSync(".");

files = files.filter(function (file) {
  var stat = fs.statSync(file);
  if (!stat.isDirectory()) {
    return;
  }
  if (/^\.+(git)?$/.test(file)) {
    return;  
  }
  return true;
});

var prefix = 'github.com/endemic-elderberry/GoHelpers'
var outfile = fs.createWriteStream('./index.go');

outfile.write(`package GoHelpers

//go:generate node generate.js
`);

return;

files.forEach(function (file) {
  outfile.write(util.format('import "%s/%s"\n', prefix, file))
});



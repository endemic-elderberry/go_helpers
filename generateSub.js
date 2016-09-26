#!/usr/bin/env node

var child_process = require("child_process");
var fs = require('fs');
var path = require('path');


process.chdir(__dirname);

var files = fs.readdirSync(__dirname);
files = files.map(function (file, i) {
  if (/^\.+/.test(file)) {
    return null;
  }
  return path.join(__dirname, file);
}).filter(function (file) {
  return file;
});

files.forEach(function (file) {
  var fstat = fs.statSync(file);
  if (!fstat || !fstat.isDirectory()) {
    return;
  }
  
  var generated = child_process.spawnSync("go",[
    "generate"
    ],{
    env:process.env,
    stdio: "inherit",
    cwd: file
  });
  if (generated.status != 0) {
    process.exit(generate.status);  
  }
});

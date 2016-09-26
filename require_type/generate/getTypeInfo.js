var jsdom = require('jsdom');

function main(){
  jsdom.env("https://golang.org/pkg/builtin/",function (err, window) {
    if (err) {
      throw err;  
    }
    var document = window.document;
    var htwos = document.querySelectorAll(".container > h2");
    htwos = [].slice.call(htwos);
    
    var typeList = [];
    htwos.forEach(function (htwo, i) {
      var textData = htwo.textContent;
      var match = /^type ([a-z].*)$/g.exec(textData);
      if (!match) {
        return;  
      }
      typeList.push(match[1]);
    });
    console.log(JSON.stringify(typeList));
  });
}

if (require.main == module) {
  main();
} else {
  var child_process = require('child_process');
  var getData = child_process.spawnSync("node", ["./getTypeInfo.js"], {
    env: process.env,
    stdio: "pipe",
    cwd: __dirname
  });
  if (getData.status !== 0) {
    throw new Error("invalid");  
  }
  var types = getData.stdout.toString();
  types = JSON.parse(types);
  module.exports = types;
}




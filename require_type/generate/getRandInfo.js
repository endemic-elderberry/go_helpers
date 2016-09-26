var jsdom = require('jsdom');

function main(){
  jsdom.env("https://golang.org/pkg/math/rand/",function (err, window) {
    if (err) {
      throw err;  
    }
    var document = window.document;
    var htwos = document.querySelectorAll(".container > h2");
    htwos = [].slice.call(htwos);
    
    var typeList = [];
    htwos.forEach(function (htwo, i) {
      var foundPre = null;
      for(var j = htwo.nextSibling; j != null; j = j.nextSibling) {
        if (String(j.tagName).toLowerCase() == "pre") {
          foundPre = j;
          break;
        }
      }
      if (!foundPre) {
        console.error("no pre found");
        return;
      }
      var textData = foundPre.textContent;
      var match = /^func ([a-zA-Z0-9]+)\(\) ([a-z0-9]+)$/g.exec(textData);
      if (!match) {
        return;  
      }
      
      typeList.push({
        generateFunc : match[1],
        primitiveType: match[2]
      });
    });
    console.log(JSON.stringify(typeList));
  });
}

if (require.main == module) {
  main();
} else {
  var child_process = require('child_process');
  var getData = child_process.spawnSync("node", ["./getRandInfo.js"], {
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




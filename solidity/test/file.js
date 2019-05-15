const fs = require('fs');

// path should be a relative path.
exports.ReadJsonFile = function (path) {
  let rawdata = fs.readFileSync(__dirname + "/" + path);
  return JSON.parse(rawdata);
};

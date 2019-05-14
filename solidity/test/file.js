const fs = require('fs');

exports.ReadJsonFile = function (path) {
  let rawdata = fs.readFileSync(path);
  return JSON.parse(rawdata);
};

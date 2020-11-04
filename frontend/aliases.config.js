const path = require('path');

function resolveSrc(_path) {
  return path.join(__dirname, _path);
}

const aliases = {
  '@src': 'src',
  '@styles': 'src/assets/styles',
  '@router': 'src',
  '@views': 'src/views',
  '@layouts': 'src/layouts',
  '@assets': 'src/assets',
  '@state': 'src/state',
  '@services': 'src/services',
  '@components': 'src/components',
};

module.exports = {
  webpack: {}
};

for (const alias in aliases) {
  module.exports.webpack[alias] = resolveSrc(aliases[alias]);
}

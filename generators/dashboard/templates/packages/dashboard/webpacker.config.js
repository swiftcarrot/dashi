const path = require('path');

const alias = ['images', 'icons', 'pages', 'layouts', 'utils'].reduce(
  (x, d) => {
    x[d] = path.resolve(__dirname, 'src', d);
    return x;
  },
  {}
);

module.exports = {
  webpack(config) {
    config.resolve = { alias };
    return config;
  },
};

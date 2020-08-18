const webpack = require('webpack');
const webpackConfig = require('../webpack');

module.exports = function () {
  const compiler = webpack(webpackConfig);

  compiler.watch({}, () => {});
};

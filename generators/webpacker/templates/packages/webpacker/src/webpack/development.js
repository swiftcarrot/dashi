const base = require('./base');

module.exports = {
  ...base,
  mode: 'development',
  output: {
    ...base.output,
    filename: 'static/[name].js',
    chunkFilename: 'static/[name].js',
    publicPath: '/',
  },
  devtool: 'cheap-module-source-map',
};

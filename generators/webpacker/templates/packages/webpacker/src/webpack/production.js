const TerserPlugin = require('terser-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const base = require('./base');

module.exports = {
  ...base,
  mode: 'production',

  output: {
    ...base.output,
    filename: 'static/[name].[chunkhash:8].js',
    chunkFilename: 'static/[name].[chunkhash:8].chunk.js',
    publicPath: '/',
  },

  optimization: {
    runtimeChunk: true,
    splitChunks: {
      chunks: 'all',
    },
    minimizer: [
      new TerserPlugin({
        parallel: true,
        cache: true,
        sourceMap: true,
        terserOptions: {
          parse: {
            ecma: 8,
          },
          compress: {
            ecma: 5,
            warnings: false,
            comparisons: false,
            drop_console: true,
          },
          mangle: {
            safari10: true,
          },
          output: {
            ecma: 5,
            comments: false,
            ascii_only: true,
          },
        },
      }),
      new OptimizeCSSAssetsPlugin({}),
    ],
  },
};

const path = require('path');
const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const WebpackAssetsManifest = require('webpack-assets-manifest');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { paths, isEnvProduction } = require('../utils/config');
const rules = require('./rules');

const alias = ['layouts', 'components', 'images', 'pages'].reduce((x, d) => {
  x[d] = path.join(paths.appSrc, d);
  return x;
}, {});

module.exports = {
  entry: {
    index: './src/index.js',
  },
  output: {
    path: paths.appBuild,
  },
  module: {
    rules: rules,
  },
  resolve: {
    alias: alias,
  },
  plugins: [
    new webpack.EnvironmentPlugin(JSON.parse(JSON.stringify(process.env))),

    new MiniCssExtractPlugin({
      filename: isEnvProduction
        ? 'static/[name].[contenthash:8].css'
        : 'static/[name].css',
      chunkFilename: isEnvProduction
        ? 'static/[name].[contenthash:8].chunk.css'
        : 'static/[name].chunk.css',
    }),

    new WebpackAssetsManifest({
      output: 'assets-manifest.json',
      entrypoints: true,
      publicPath: true,
      writeToDisk: true,
    }),

    new HtmlWebpackPlugin({
      filename: 'index.html',
      minify: true,
      template: paths.appHtml,
    }),
  ],
};

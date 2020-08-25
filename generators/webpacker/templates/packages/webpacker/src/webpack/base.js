const webpack = require('webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const WebpackAssetsManifest = require('webpack-assets-manifest');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { paths, isEnvProduction } = require('../utils/config');
const rules = require('./rules');

module.exports = {
  entry: {
    index: './src/index.js',
  },
  output: {
    path: paths.appBuild,
    publicPath: paths.publicUrlOrPath,
  },
  module: {
    rules: rules,
  },
  resolve: {
    alias: {
      src: paths.appSrc,
    },
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

    new HtmlWebpackPlugin(
      Object.assign(
        {},
        {
          inject: true,
          template: paths.appHtml,
        },
        isEnvProduction
          ? {
              minify: {
                removeComments: true,
                collapseWhitespace: true,
                removeRedundantAttributes: true,
                useShortDoctype: true,
                removeEmptyAttributes: true,
                removeStyleLinkTypeAttributes: true,
                keepClosingSlash: true,
                minifyJS: true,
                minifyCSS: true,
                minifyURLs: true,
              },
            }
          : undefined
      )
    ),
  ],
};

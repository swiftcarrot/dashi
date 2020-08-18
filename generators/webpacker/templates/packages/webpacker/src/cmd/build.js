const fs = require('fs-extra');
const webpack = require('webpack');
const { paths } = require('../utils/config');
const webpackConfig = require('../webpack');

module.exports = function () {
  if (fs.pathExistsSync(paths.appPublic)) {
    fs.copySync(paths.appPublic, paths.appBuild);
  }

  webpack(webpackConfig, (err, stats) => {
    if (err || stats.hasErrors()) {
      console.error(err);
    }

    console.log(stats.toString({ chunks: false, colors: true }));
  });
};

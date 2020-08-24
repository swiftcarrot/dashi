const glob = require('glob');
const rules = require('webpacker/src/webpack/rules');

const alias = glob
  .sync('packages/**/package.json', {
    ignore: '**/node_modules/**/package.json',
  })
  .reduce((x, file) => {
    const { name } = require(file.replace('packages/', ''));
    x[name] = `${name}/src`;
    return x;
  }, {});

const webpack = (webpackConfig = {}, options = { lessOptions: {} }) => {
  const { module = {} } = webpackConfig;
  return {
    ...webpackConfig,
    module: {
      ...module,
      rules: [...(module.rules || []), ...rules],
    },
  };
};

const babel = (config = {}) => {
  return {
    presets: ['module:babel-preset', ...config.presets],
  };
};

module.exports = { webpack, babel };

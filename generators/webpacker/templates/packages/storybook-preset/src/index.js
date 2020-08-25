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
  const { module = {}, resolve = {} } = webpackConfig;
  return {
    ...webpackConfig,
    resolve: {
      ...resolve,
      alias: {
        ...resolve.alias,
        ...alias,
      },
    },
    module: {
      ...module,
      rules: [...(module.rules || []), ...rules],
    },
  };
};

// Don't use Storybook's default Babel config.
const babelDefault = (config = {}) => {
  return {
    presets: [],
    plugins: [],
  };
};

module.exports = { webpack, babelDefault };

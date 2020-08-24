const glob = require('glob');

const alias = glob
  .sync('packages/**/package.json', {
    ignore: '**/node_modules/**/package.json',
  })
  .reduce((x, file) => {
    const { name } = require(`../${file}`);
    x[name] = `${name}/src`;
    return x;
  }, {});

module.exports = {
  stories: ['../stories/**/*.stories.js'],
  addons: ['@storybook/addon-actions', '@storybook/addon-links'],
  webpackFinal: async (config) => {
    config.resolve.alias = alias;
    return config;
  },
};

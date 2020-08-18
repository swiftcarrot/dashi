const fs = require('fs');
const path = require('path');

const appDirectory = fs.realpathSync(process.cwd());
const resolveApp = (relativePath) => path.resolve(appDirectory, relativePath);

const isEnvDevelopment = process.env.NODE_ENV === 'development';
const isEnvProduction = process.env.NODE_ENV === 'production';
const shouldUseSourceMap = process.env.GENERATE_SOURCEMAP !== 'false';

const paths = {
  appSrc: resolveApp('src'),
  appBuild: resolveApp('build'),
  appPublic: resolveApp('public'),
  appHtml: resolveApp('public/index.html'),
  publicUrlOrPath:
    require(resolveApp('package.json')).homepage ||
    process.env.PUBLIC_URL ||
    '/',
};

module.exports = {
  isEnvDevelopment,
  isEnvProduction,
  shouldUseSourceMap,
  paths,
};

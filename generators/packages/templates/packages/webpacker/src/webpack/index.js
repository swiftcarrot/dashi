const { isEnvProduction } = require('../utils/config');

module.exports = isEnvProduction
  ? require('./production')
  : require('./development');

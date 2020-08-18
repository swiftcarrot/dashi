const { parse } = require('url');
const webpack = require('webpack');
const Server = require('webpack-dev-server/lib/Server');
const { paths } = require('../utils/config');
const webpackConfig = require('../webpack');

module.exports = (cmd) => {
  console.log('listen', cmd.args[0]);
  const listen = cmd.args[0];

  const compiler = webpack(webpackConfig);
  const server = new Server(compiler, {
    contentBase: paths.appPublic,
    historyApiFallback: true,
  });
  const url = parse(listen);
  const { port, hostname } = url;

  server.listen(port, hostname, () => {
    console.log(`Starting server on ${hostname}:${port}`);
  });
};

#!/usr/bin/env node
const { Command } = require('commander');
const serve = require('./cmd/serve');
const build = require('./cmd/build');
const watch = require('./cmd/watch');

const program = new Command();

program.version(require('../package.json').version);

program
  .command('serve')
  .description('Start webpack-dev-server')
  .option(
    '-l',
    '--listen <listen>',
    'server listening address http://0.0.0.0:3000',
    null,
    'http://0.0.0.0:3000'
  )
  .action(serve);
program.command('build').description('Webpack build').action(build);
program.command('watch').description('Webpack watch').action(watch);

program.parse(process.argv);

// console.log('webpacker');

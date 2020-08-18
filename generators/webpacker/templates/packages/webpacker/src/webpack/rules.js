const getStyleLoaders = require('../utils/getStyleLoaders');
const getCSSModuleLocalIdent = require('../utils/getCSSModuleLocalIdent');
const {
  isEnvProduction,
  isEnvDevelopment,
  shouldUseSourceMap,
} = require('../utils/config');

const cssRegex = /\.css$/;
const cssModuleRegex = /\.module\.css$/;
const sassRegex = /\.(scss|sass)$/;
const sassModuleRegex = /\.module\.(scss|sass)$/;

module.exports = [
  {
    oneOf: [
      {
        test: /icons\/.*.\.svg$/,
        use: [{ loader: require.resolve('@svgr/webpack') }],
      },

      {
        test: /\.(jpg|jpeg|png|gif|svg|eot|ttf|woff|woff2)$/i,
        use: [
          {
            loader: require.resolve('file-loader'),
            options: {
              name: 'static/[name].[hash:8].[ext]',
              esModule: false,
            },
          },
        ],
      },
    ],
  },

  {
    test: /\.js$/,
    exclude: /node_modules/,
    loader: require.resolve('babel-loader'),
    options: {
      presets: [require.resolve('babel-preset')],
      cacheDirectory: true,
    },
  },

  {
    test: /\.ya?ml$/,
    use: require.resolve('@swiftcarrot/yaml-loader'),
  },

  {
    test: cssRegex,
    exclude: cssModuleRegex,
    use: getStyleLoaders({
      importLoaders: 1,
      sourceMap: isEnvProduction ? shouldUseSourceMap : isEnvDevelopment,
    }),
    sideEffects: true,
  },

  {
    test: cssModuleRegex,
    use: getStyleLoaders({
      importLoaders: 1,
      sourceMap: isEnvProduction ? shouldUseSourceMap : isEnvDevelopment,
      modules: {
        getLocalIdent: getCSSModuleLocalIdent,
      },
    }),
  },

  {
    test: sassRegex,
    exclude: sassModuleRegex,
    use: getStyleLoaders(
      {
        importLoaders: 3,
        sourceMap: isEnvProduction ? shouldUseSourceMap : isEnvDevelopment,
      },
      'sass-loader'
    ),
    sideEffects: true,
  },

  {
    test: sassModuleRegex,
    use: getStyleLoaders(
      {
        importLoaders: 3,
        sourceMap: isEnvProduction ? shouldUseSourceMap : isEnvDevelopment,
        modules: {
          getLocalIdent: getCSSModuleLocalIdent,
        },
      },
      'sass-loader'
    ),
  },
];

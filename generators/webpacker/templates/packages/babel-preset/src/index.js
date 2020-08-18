module.exports = function(api, { production }) {
  const prod = production || process.env.NODE_ENV === 'production';

  return {
    presets: [
      [
        require('@babel/preset-env'),
        {
          loose: true,
          targets: {
            ie: 9,
          },
        },
      ],
      require('@babel/preset-react'),
      require('@babel/preset-flow'),
    ],
    plugins: [
      require('@babel/plugin-transform-runtime'),
      require('@babel/plugin-proposal-export-default-from'),
      require('@babel/plugin-syntax-dynamic-import'),
      [require('@babel/plugin-proposal-class-properties'), { loose: true }],
      prod && [
        require('babel-plugin-transform-react-remove-prop-types'),
        { removeImport: true },
      ],
    ].filter(Boolean),
  };
};

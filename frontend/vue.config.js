const webpack = require('webpack');
const appConfig = require('./src/app.config');

module.exports = {
  configureWebpack: {
    name: appConfig.title,
    resolve: {
      alias: require('./aliases.config').webpack
    },
    plugins: [new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/)]
  },
  css: {
    sourceMap: true
  },
  devServer: {
    proxy: "http://localhost:9000/"
  }
};

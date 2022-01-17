
const path = require('path')

function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
    lintOnSave: process.env.NODE_ENV === 'development',
    productionSourceMap: false,
    devServer: {
        open: false,
        port: 4000
    },
    publicPath: "/",
    configureWebpack: {
        module: {
            rules: [
                {
                    test: /\.mjs$/,
                    include: /node_modules/,
                    type: "javascript/auto"
                },
            ],
        },
    },
    chainWebpack: config => {
        config.module.rule('svg').exclude.add(resolve('src/icons')).end()
        config.module.rule('icons').test(/\.svg$/).include.add(resolve('src/icons')).end().use('svg-sprite-loader').loader('svg-sprite-loader').options({
            symbolId: 'icon-[name]'
        }).end()
    }

}
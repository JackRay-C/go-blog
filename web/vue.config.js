module.exports = {
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
    }
}
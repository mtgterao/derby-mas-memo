var path = require('path');

module.exports = {
    entry: path.join(__dirname, 'static/js/app.js'),
    output: {
        path: path.join(__dirname, 'static/dist'),
        filename: 'bundle.js'
    },
    cache: true,
    module: {
        loaders: [{
            test: /\.css$/, loader: 'style-loader!css-loader'
        }, {
            test: /\.js[x]?$/,
            exclude: /node_modules/,
            loader: 'babel',
            query: {
                presets: ['react', 'es2015']
            }
        }, {
            test: /\.js$/,
            exclude: /node_modules/,
            loader: 'eslint-loader'
        }]
    },
    eslint: {
        configFile: "./.eslintrc"
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    },
    devtool: 'source-map'
}

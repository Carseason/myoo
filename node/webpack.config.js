const path = require('path');
const webpack = require('webpack');//引入webpack
const VueLoaderPlugin = require('vue-loader/lib/plugin')
const TerserPlugin = require('terser-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = {
    mode: 'production',//开发模式development,生产模式production
    /*************Myoo**************** */
    entry: ["@babel/polyfill", "/go/mod/myoo/node/app.js"],
    output: { filename: 'app.js', path: "/go/mod/myoo/src/js" },
    devServer: {
        contentBase: "/go/mod/myoo/src/",
        compress: true,
        open: true,
        openPage: 'admin.html#/admin',
        hot: true,
        port: 8000
    },
    plugins: [    // 引入插件
        new VueLoaderPlugin(),
        new webpack.HotModuleReplacementPlugin(),
        new MiniCssExtractPlugin({
            filename: "style.css",
            chunkFilename: '[id].css',
            ignoreOrder: false,
        }),
    ],
    resolve: {
        modules: ["/node_modules"],    //绝对路径引用npm包
        alias: {
            "@": path.join(__dirname, "./"),
            "vue$": "vue/dist/vue.min.js",  //修改vue包导入路径
        }
    },
    optimization: {
        minimize: true,
        minimizer: [
            new TerserPlugin({
                extractComments: false,
                cache: '/node_modules/.cache/terser-webpack-plugin',
                //cache: false,

            }),
        ],
    },
    module: {    //用于配置第三方模块
        rules: [//所有第三方模块的匹配规则
            { test: /\.(js|jsx)$/, exclude: /(node_modules|bower_components)/, use: { loader: 'babel-loader', options: { presets: ['@babel/preset-env'] } } },
            { test: /\.css$/, use: [MiniCssExtractPlugin.loader, 'css-loader'], },
            // { test: /\.css$/, use: ['style-loader', 'css-loader?modules'] }, //配置处理第三方文件的规则,?modules启用模块化
            { test: /\.(png|jpg|jpeg|gif|eot|ttf|woff|woff2|svg|svgz)$/, loader: 'file-loader', },
            { test: /\.vue$/, use: 'vue-loader' }    //处理.vue文件的渲染
        ]
    },

};

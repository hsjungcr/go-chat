var webpack = require('webpack');
var path = require('path');

var BUILD_DIR = path.resolve(__dirname, 'public');
var APP_DIR = path.resolve(__dirname, 'src/app');

var config = {
		entry: APP_DIR + '/index.jsx',
		output: {
			path: BUILD_DIR,
			filename: 'client.min.js'
		},
		resolve: {
			extensions: ['.js', '.jsx']
		},
		module: {
			loaders: [{
					test: /\.jsx?$/,
					include: APP_DIR,
					loader: 'babel-loader',
					query: {
						presets: ['es2017', 'react']}
					}
				]
			}
		};

		module.exports = config;
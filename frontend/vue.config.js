module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:8001',
                changeOrigin: true,
                secure: false,
                ws: true,
                pathRewrite:{
                    "^/api":"/api"
                }
            }
        },
    },
};
const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = (app) => {
  app.use(
    '/style.css',
    createProxyMiddleware({
      target: 'http://localhost:3001/style.css',
      changeOrigin: true,
    }),
  );
  app.use(
    '/favicon.ico',
    createProxyMiddleware({
      target: 'http://localhost:3001/favicon.ico',
      changeOrigin: true,
    }),
  );
};

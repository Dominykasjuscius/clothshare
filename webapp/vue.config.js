module.exports = {
    devServer: {
      "proxy": {
        "/api": {
        "target": 'http://localhost:8088',
        "secure": false
        }
      }
    }
  };
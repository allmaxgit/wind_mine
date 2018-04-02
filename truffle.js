require('babel-register');
require('babel-polyfill');

module.exports = {
    migrations_directory: './migrations',
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
    networks: {
        development: {
            host: 'localhost',
            port: 8545,
            network_id: '*' // Match any network id
        }
    }
};

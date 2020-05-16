'use strict';

const path = require('path');
const dotenv = require('dotenv');

if (process.env.APP_ENV === 'test') {
    dotenv.config({
        path: path.resolve(process.cwd(), '.env.test'),
    });
} else {
    dotenv.config();
}

const express = require('express');
const mongose = require('mongoose');
const cors = require('cors');
const bodyParser = require('body-parser');
const dbConfig = require('./database/db');
const api = require('./routes/auth.routes');

mongose.Promise = global.Promise;
mongose.connect(dbConfig.db, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
}).then(() => {
    console.log('Database connected');
}, err => {
    console.log('Error while connect to database: ' + err)
});
mongose.set('useCreateIndex', true);

const app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true,
}));
app.use(cors());
app.use('/public', express.static('public'));
app.use('/api', api);

const port = process.env.APP_PORT || 4000;
const server = app.listen(port, () => {
    console.log('Connected to port: ' + port)
});

app.use((req, res, next) => {
    setImmediate(() => {
        next(new Error('Something went wrong'));
    });
});

app.use((err, req, res, next) => {
    console.log(err);
    if (!err.statusCode) {
        err.statusCode = 500
    }
    return res.status(err.statusCode).send({
        message: 'Internal server error.',
    })
});

module.exports = server;

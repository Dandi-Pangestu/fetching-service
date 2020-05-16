'use strict';

const mongoose = require('mongoose');
const Schema = mongoose.Schema;
const uniqueValidator = require('mongoose-unique-validator');

let userSchema = new Schema({
    name: {
        type: String,
        unique: true,
    },
    phone: {
        type: String,
    },
    role: {
        type: String,
    },
    password: {
        type: String,
    },
}, {
    collections: 'users',
});

userSchema.plugin(uniqueValidator, { message: 'User name already in use.' });
module.exports = mongoose.model('User', userSchema);

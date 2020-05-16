'use strict';

const express = require('express');
const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');
const router = express.Router();
const userSchema = require('../models/user');
const { check, validationResult } = require('express-validator');
const authorize = require('../middlewares/auth');

router.post('/register',
    [
        check('name')
            .not()
            .isEmpty(),
        check('phone')
            .not()
            .isEmpty(),
        check('role')
            .not()
            .isEmpty(),
    ], (req, res, next) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
        return res.status(422).jsonp(errors.array());
    }

    let randPassword = Math.random().toString(36).substring(2, 6);
    bcrypt.hash(randPassword, 10).then((hash) => {
        let user = new userSchema({
            name: req.body.name,
            phone: req.body.phone,
            role: req.body.role,
            password: hash,
        });
        user.save().then((response) => {
            let jwtToken = jwt.sign({
                id: response._id,
                name: response.name,
                phone: response.phone,
                role: response.role,
            }, process.env.JWT_SECRET_KEY, {
                expiresIn: '1h',
            });
            res.status(200).json({
                token: jwtToken,
                password: randPassword,
            });
        }).catch((err) => {
            res.status(500).json({
                message: 'Internal server error.',
                error: err,
            });
        });
    });
});

router.post('/login',
    [
        check('phone')
            .not()
            .isEmpty(),
        check('password')
            .not()
            .isEmpty(),
    ], (req, res, next) => {
    let errors = validationResult(req);
    if (!errors.isEmpty()) {
        return res.status(422).jsonp(errors.array());
    }

    let user;
    userSchema.findOne({
        phone: req.body.phone,
    }).then(u => {
        if (!u) {
            return res.status(401).json({
                message: 'Unauthenticated.',
            })
        }

        user = u;
        return bcrypt.compare(req.body.password, user.password);
    }).then(response => {
        if (!response) {
            return res.status(401).json({
                message: 'Unauthenticated.',
            })
        }

        let jwtToken = jwt.sign({
            id: user._id,
            name: user.name,
            phone: user.phone,
            role: user.role,
        }, process.env.JWT_SECRET_KEY, {
            expiresIn: '1h',
        });

        res.status(200).json({
            token: jwtToken,
        });
    }).catch(err => {
        return res.status(401).json({
            message: 'Unauthenticated.',
        })
    });
});

router.route('/profile').get(authorize, (req, res, next) => {
    res.status(200).json({
        name: req.authUser.name,
        phone: req.authUser.phone,
        role: req.authUser.role,
    });
});

module.exports = router;

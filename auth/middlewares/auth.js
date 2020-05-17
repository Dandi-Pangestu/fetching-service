'use strict';

const jwt = require('jsonwebtoken');

module.exports = (req, res, next) => {
    try {
        const token = req.headers.authorization.split(' ')[1];
        let decoded = jwt.verify(token, process.env.JWT_SECRET_KEY);
        req.authUser = {
            id: decoded.id,
            name: decoded.name,
            phone: decoded.phone,
            role: decoded.role,
            timestamp: decoded.timestamp,
        };
        next();
    } catch (e) {
        res.status(401).json({
            message: 'Unauthenticated.',
        });
    }
};

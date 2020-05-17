'use strict';

const app = require('../server');
const userSchema = require('../models/user');
const chai = require('chai');
const chaiHttp = require('chai-http');
chai.use(chaiHttp);
const should = chai.should();
const expect = chai.expect;
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');

before(() => {
    console.log('Testing started...');
    truncate();
});

after(() => {
    truncate();
    console.log('Testing completed...');
});

beforeEach(() => {
    truncate();
});

afterEach(() => {
    truncate();
});

function createUser() {
    let user;

    return bcrypt.hash('1234', 10).then((hash) => {
        let user = new userSchema({
            name: 'John',
            phone: '08756756757',
            role: 'admin',
            password: hash,
            timestamp: Date.now().toString(),
        });
        user.save().then(response => {
            user = response;
        });

        return new Promise(function (resolve, reject) {
            resolve(user);
        });
    });
}

function truncate() {
    userSchema.deleteMany({}, err => {
        if (err) {
            console.log('Error while truncate user: ' + err);
        }
    });
}

describe('Register', () => {
    let url = '/api/register';

    describe('Validation error expect', () => {
        it('Body request is required', (done) => {
            chai
                .request(app)
                .post(url)
                .set('Content-Type', 'application/json')
                .send({})
                .end((err, res) => {
                    res.should.have.status(422);
                    expect(res.body).to.deep.equal([
                        {
                            "msg": "Invalid value",
                            "param": "name",
                            "location": "body"
                        },
                        {
                            "msg": "Invalid value",
                            "param": "phone",
                            "location": "body"
                        },
                        {
                            "msg": "Invalid value",
                            "param": "role",
                            "location": "body"
                        }
                    ]);
                    done();
                });
        });

        it('User name already in use', (done) => {
            createUser().then((user) => {
                chai
                    .request(app)
                    .post(url)
                    .set('Content-Type', 'application/json')
                    .send({
                        name: 'John',
                        phone: '123',
                        role: 'admin',
                    })
                    .end((err, res) => {
                        res.should.have.status(500);
                        done();
                    });
            });
        });
    });

    it('Register successful', (done) => {
        chai
            .request(app)
            .post(url)
            .set('Content-Type', 'application/json')
            .send({
                name: 'John',
                phone: '123',
                role: 'admin',
            })
            .end((err, res) => {
                res.should.have.status(200);
                done();
            });
    });
});

describe('Login', () => {
    let url = '/api/login';

    describe('Validation error expect', () => {
        it('Body request is required', (done) => {
            chai
                .request(app)
                .post(url)
                .set('Content-Type', 'application/json')
                .send({})
                .end((err, res) => {
                    res.should.have.status(422);
                    expect(res.body).to.deep.equal([
                        {
                            "msg": "Invalid value",
                            "param": "phone",
                            "location": "body"
                        },
                        {
                            "msg": "Invalid value",
                            "param": "password",
                            "location": "body"
                        }
                    ]);
                    done();
                });
        });

        it('Unauthenticated', (done) => {
            chai
                .request(app)
                .post(url)
                .set('Content-Type', 'application/json')
                .send({
                    phone: '08756756757',
                    password: '123',
                })
                .end((err, res) => {
                    res.should.have.status(401);
                    expect(res.body).to.deep.equal({
                        message: 'Unauthenticated.',
                    });
                    done();
                });
        });
    });

    it('Login successful', (done) => {
        createUser().then((user) => {
            chai
                .request(app)
                .post(url)
                .set('Content-Type', 'application/json')
                .send({
                    phone: '08756756757',
                    password: '1234',
                })
                .end((err, res) => {
                    res.should.have.status(200);
                    done();
                });
        });
    });
});

describe('Get Profile', () => {
    let url = '/api/profile';

    it('Should return 401', (done) => {
        createUser().then((user) => {
            chai
                .request(app)
                .get(url)
                .set('Authorization', '')
                .end((err, res) => {
                    res.should.have.status(401);
                    done();
                });
        })
    });

    it('Get profile successful', (done) => {
        createUser().then((user) => {
            let jwtToken = jwt.sign({
                id: user._id,
                name: user.name,
                phone: user.phone,
                role: user.role,
                timestamp: user.timestamp,
            }, process.env.JWT_SECRET_KEY, {
                expiresIn: '1h',
            });

            chai
                .request(app)
                .get(url)
                .set('Authorization', 'Bearer ' + jwtToken)
                .end((err, res) => {
                    res.should.have.status(200);
                    expect(res.body).to.deep.equal({
                        name: user.name,
                        phone: user.phone,
                        role: user.role,
                        timestamp: user.timestamp,
                    });
                    done();
                });
        })
    });
});

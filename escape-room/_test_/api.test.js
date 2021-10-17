const app = require('./test-server/server.js');
const supertest = require('supertest');
const request = supertest(app);

describe('testing to see if jest works', () => {
    test('test validation (happy Path)', () => {
        expect(1).toBe(1);
    })
    test('test validation (sad Path) ', () => {
        expect(1).toBe(2);
    })
})
describe('API Testing (local server)', () => {
    test('API validation (happy Path)', async done => {
        const response = await request.get('/api')
        expect(response.status).toBe(200);
        done();
    })
    test('API validation (sad Path) ', () => {
        
    })
})
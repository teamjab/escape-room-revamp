const app = require('./server.js');

const port = 3000;

app.get('/', (request,response) => {
    response.send('Test server');
})

app.listen(port, () => {
    console.log(`app is up on port:${port}`);
});
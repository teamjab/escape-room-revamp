const app = require('./server.js');

const port = 3000;

app.get('/api', async (request,response) => {
    response.json({message:'API test route'});
});

app.listen(port, () => {
    console.log(`app is up on port:${port}`);
});
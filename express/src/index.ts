import express from 'express'

const app = express();

app.get('/', (_req, res) => {
    res.json({ hello: "world" })
});


app.listen(3001, () => console.log('Express is listening on 3001'));

const path = require('path')
const express = require('express');
const bodyParser = require('body-parser');
const notifier = require('node-notifier');
const app = express();
const port = process.env.PORT || 9000;

app.use(bodyParser.json());

app.get("/health", (req, res) => {
    res.status(200).send();
})

app.post("/notify", (req, res) => {
    notify(req.body, reply => {
        res.send(reply)
    })
})

app.listen(port, () => console.log("server is up and running on port", port));

const notify = ({ title, message }, cb) => {
    console.log(title,message)
    notifier.notify(
        {
            title: title || "Unknown message",
            message: message || "Unknown message",
            sound: true,
            wait: true,
            reply: true,
            icon: path.join(__dirname,"logo.png"),
            closeLabel: "Completed?",
            timeout: 15
        },
        (err, response) => {
            console.log(response)
            cb(response)
        })
}
const http = require("http");
const https = require("https");
const fs = require("fs");
const express = require("express");
const path = require("path");

const app = express();

const host = "0.0.0.0";
const port = 8001;

var privateKey = fs.readFileSync("../../certification/server-key.pem");
var certificate = fs.readFileSync("../../certification/server-cert.pem");
var credentials = { key: privateKey, cert: certificate };

app.use(express.static(path.join(__dirname, "public")));
app.use("/dist", express.static(path.join(__dirname, "dist")));

app.get("/", (request, response) => {
  response.sendFile(__dirname + "/index.html");
});

var httpServer = http.createServer(app);
var httpsServer = https.createServer(credentials, app);

httpServer.listen(port, host);
console.log(`Running on http://${host}:${port}`);

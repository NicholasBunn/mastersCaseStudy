const http = require("http");
const fs = require("fs");
const express = require("express");
const path = require("path");

const app = express();

const host = "localhost";
const port = 8000;

app.use(express.static(path.join(__dirname, "public")));
app.use("/dist", express.static(path.join(__dirname, "dist")));

app.get("/", (request, response) => {
  response.sendFile(__dirname + "/index.html");
});

var privateKey = fs.readFileSync("../../certification/server-key.pem");
var serverCert = fs.readFileSync("../../certification/server-cert.pem");

http
  .createServer({ key: privateKey, cert: serverCert }, app)
  .listen(port, host);

const express = require("express");
const cors = require("cors");

const app = express();
const port = 8000;

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to nimish server");
});

app.post("/", (req, res) => {
  res.status(201).send({ message: req.body });
});
app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});
app.listen(port, () => {
  console.log("Listening on port: ", port);
});

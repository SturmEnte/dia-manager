require("dotenv/config");
const express = require("express");
const path = require("path");

const app = express();

app.use(express.static(path.join(__dirname, "public")));

app.get("/", (req, res) => {
	res.sendFile(path.join(__dirname, "public", "main/main.html"));
});

app.listen(process.env.PORT, () => {
	console.log("Listening on port " + process.env.PORT);
});

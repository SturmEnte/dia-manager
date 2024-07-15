require("dotenv/config");
const express = require("express");
const cookieParser = require("cookie-parser");
const path = require("path");

const WHITE_LIST = ["api", "login", "signup"];

const app = express();

app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));

app.all("*", (req, res, next) => {
	for (whiteListWord of WHITE_LIST) {
		if (req.url.includes(whiteListWord)) {
			next();
			return;
		}
	}

	if (req.cookies["loggedIn"] == undefined) {
		res.redirect("/login");
		return;
	}

	next();
});

app.get("/", (req, res) => {
	res.sendFile(path.join(__dirname, "public", "main/main.html"));
});

app.get("/login", (req, res) => {
	res.sendFile(path.join(__dirname, "public", "login/login.html"));
});

app.listen(process.env.PORT, () => {
	console.log("Listening on port " + process.env.PORT);
});

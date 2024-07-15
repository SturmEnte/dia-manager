require("dotenv/config");
const express = require("express");
const cookieParser = require("cookie-parser");
const path = require("path");
const mysql = require("mysql2/promise");

const setupDatabase = require("./util/setupDatabase.js");

const WHITE_LIST = ["api", "login", "signup"];

(async () => {
	// Connect to database
	let connection;
	try {
		connection = await mysql.createConnection({
			host: process.env.MYSQL_HOST,
			user: process.env.MYSQL_USER,
			password: process.env.MYSQL_PASSWORD,
			database: process.env.MYSQL_DATABASE,
		});
	} catch (error) {
		console.error("Error while connecting to database");
		throw error;
	}

	await setupDatabase(connection);

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

	app.get("/signup", (req, res) => {
		res.sendFile(path.join(__dirname, "public", "signup/signup.html"));
	});

	app.listen(process.env.PORT, () => {
		console.log("Listening on port " + process.env.PORT);
	});
})();

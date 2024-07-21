import "dotenv/config";
import express from "express";
import cookieParser from "cookie-parser";
import path from "path";
import pg from "pg";

import setupDatabase from "./util/setupDatabase";
import getLanguageFilePath from "./util/getLanguageFilePath";

import signup from "./routes/api/auth/signup";
import login from "./routes/api/auth/login";

const config = require("../configs/config.json");

const WHITE_LIST = ["api", "login", "signup"];

(async () => {
	// Connect to database
	const client = new pg.Client({
		host: process.env.PG_HOST,
		port: Number(process.env.PG_PORT),
		user: process.env.PG_USER,
		password: process.env.PG_PASSWORD,
		database: process.env.PG_DATABASE,
	});

	await client.connect((err) => {
		if (err) {
			console.error("Error while connecting to database");
			throw err;
		}

		console.log("Connected to database");
	});

	await setupDatabase(client);

	const app = express();

	app.use(express.static(path.join(__dirname, config.publicLocation)));

	app.use(cookieParser());

	app.all("*", (req, res, next) => {
		for (let whiteListWord of WHITE_LIST) {
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

	app.use("/api/auth/", signup(client));
	app.use("/api/auth/", login(client));

	// app.get("/", (req, res) => {
	// 	res.sendFile(getLanguageFilePath(path.join(__dirname, "public", "main/main.html"), req.cookies["language"]));
	// });

	// app.get("/login", (req, res) => {
	// 	res.sendFile(getLanguageFilePath(path.join(__dirname, "public", "login/login.html"), req.cookies["language"]));
	// });

	// app.get("/signup", (req, res) => {
	// 	res.sendFile(getLanguageFilePath(path.join(__dirname, "public", "signup/signup.html"), req.cookies["language"]));
	// });

	app.listen(process.env.PORT, () => {
		console.log("Listening on port " + process.env.PORT);
	});
})();
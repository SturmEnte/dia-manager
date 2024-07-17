import { Router, json } from "express";
import { Client } from "pg";

import { hashString } from "../../util/hashing";

let client: Client | undefined;

const router = Router();

router.use(json());

router.post("/login", async (req, res) => {
	if (!client) {
		res.status(500).json({ error: "Failed database connection" });
		return;
	}

	if (!req.body.username) {
		res.status(400).json({ error: "Username required" });
		return;
	}

	if (!req.body.password) {
		res.status(400).json({ error: "Password required" });
		return;
	}

	const username: string = req.body.username;
	const password: string = req.body.password;

	if ((await client.query(`SELECT COUNT(id) FROM users WHERE LOWER(username) = LOWER('${username}')`)).rows[0].count < 1) {
		res.status(400).json({ error: "No user with that username exists" });
		return;
	}

	try {
		// Check if the entered password matches the password stored in the database
		const savedHashedPassword = (await client.query(`SELECT password FROM users WHERE username = '${username}'`)).rows[0].password;
		const hashedPassword = hashString(password);

		if (savedHashedPassword != hashedPassword) {
			res.status(400).json({ error: "The password is incorrect" });
			return;
		}

		// Create access token
	} catch (error) {
		if (error) {
			console.error("Error while creating user");
			console.error(error);
			res.status(500).json({ error: "Error while creating user" });
			return;
		}
	}
});

export default (newClient: Client): Router => {
	client = newClient;
	return router;
};

import { Router, json } from "express";
import { Client } from "pg";

import formatDate from "../../util/formatDate";

let client: Client | undefined;

const router = Router();

router.use(json());

router.post("/signup", async (req, res) => {
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

	const username = req.body.username;
	const password = req.body.password;

	if ((await client.query(`SELECT COUNT(id) FROM users WHERE LOWER(username) = LOWER('${username}')`)).rows[0].count > 0) {
		res.status(400).json({ error: "A user with that username already exists" });
		return;
	}

	try {
		await client.query(`INSERT INTO users (username, password, created) VALUES ('${username}', '${password}', '${formatDate(new Date())}')`);
		res.json({ message: "Created user" });
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

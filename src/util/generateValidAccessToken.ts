import { Client } from "pg";
import { randomBytes } from "crypto";

const tablesConfig = require("../../configs/tables.json");

// Generate a random token and check if it is already in use
// If an error occurs during the genration or checking, then "undefined" is returned else a valid token is returned
export default async function (client: Client): Promise<string | undefined> {
	let token: string;

	try {
		do {
			token = randomBytes(tablesConfig.users.tokenLength).toString();
		} while (await client.query(`SELECT COUNT(id) FROM access_tokens WHERE token = '${token}'`));
	} catch (error) {
		console.error("Error while generating access token");
		console.error(error);
		return undefined;
	}

	return token;
}

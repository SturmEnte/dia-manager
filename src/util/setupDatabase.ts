import { Client } from "pg";

const tablesConfig = require("../../configs/tables.json");

// Check if tables exist and create them if they don't
export default async function (client: Client) {
	// Users table
	try {
		await client.query(
			`CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(${tablesConfig.users.maxUsernameLength}) NOT NULL UNIQUE,
				password VARCHAR(${tablesConfig.users.maxPasswordLength}) NOT NULL,
				created DATE NOT NULL
			);`
		);
	} catch (error) {
		console.error("Error while creating users table");
		throw error;
	}
}

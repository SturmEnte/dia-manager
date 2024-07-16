const tablesConfig = require("../../configs/tables.json");

// Check if tables exist and create them if they don't
export default async function (connection) {
	// Users table
	try {
		await connection.query(
			"CREATE TABLE IF NOT EXISTS  `users` (`id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(" +
				tablesConfig.users.maxUsernameLength +
				") NOT NULL , `password` VARCHAR(" +
				tablesConfig.users.maxPasswordLength +
				") NOT NULL , `create` DATE NOT NULL , PRIMARY KEY (`id`), UNIQUE (`username`)) ENGINE = InnoDB;"
		);
	} catch (error) {
		console.error("Error while creating users table");
		throw error;
	}
}

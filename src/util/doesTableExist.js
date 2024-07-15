module.exports = async function (connection, tableName) {
	try {
		const [rows] = await connection.query(`
        SELECT * FROM information_schema.TABLES
        WHERE TABLE_NAME = '${tableName}'
      `);

		return rows.length > 0; // True if table exists, False otherwise
	} catch (err) {
		console.error("Error:", err);
		return false; // Handle potential errors and return false
	}
};

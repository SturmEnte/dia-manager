export default function (date: Date) {
	const offset = 0; // Offset in hours for +0 timezone

	// Get individual components
	const year = date.getUTCFullYear();
	const month = String(date.getUTCMonth() + 1).padStart(2, "0"); // Pad month for single digit
	const day = String(date.getUTCDate()).padStart(2, "0");
	const hours = String(date.getUTCHours()).padStart(2, "0");
	const minutes = String(date.getUTCMinutes()).padStart(2, "0");
	const seconds = String(date.getUTCSeconds()).padStart(2, "0");

	// Build the string with offset
	return `${year}-${month}-${day} ${hours}:${minutes}:${seconds} ${offset >= 0 ? "+" : "-"}${Math.abs(offset)}:00`;
}

import { createHash, BinaryToTextEncoding } from "crypto";

// Todo: Add salt
export function hashString(string: string, algorithm: string = "sha256", encoding: BinaryToTextEncoding = "hex"): string {
	const hash = createHash(algorithm);
	hash.update(string);
	return hash.digest(encoding);
}

// Checks if the given string is the same as the hashed string
// It hashes the given string witht the algorithm and encoding and then checks if the hashed strings are the same
export function isStringMatchingHashedString(string: string, hashedString: string, algorithm: string = "sha256", encoding: BinaryToTextEncoding = "hex"): boolean {
	if (hashString(string, algorithm, encoding) === hashedString) return true;
	return false;
}

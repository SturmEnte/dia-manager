import { createHash, BinaryToTextEncoding } from "crypto";

// Todo: Add salt
export function hashString(string: string, algorithm: string = "sha256", encoding: BinaryToTextEncoding = "hex") {
	const hash = createHash(algorithm);
	hash.update(string);
	return hash.digest(encoding);
}

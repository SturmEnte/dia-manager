import { existsSync } from "fs";

export default function (path: string, language: string): string {
	const languagePath = path.replace("html", language + ".html");

	console.log(languagePath);

	if (existsSync(languagePath)) {
		return languagePath;
	}

	return path;
}

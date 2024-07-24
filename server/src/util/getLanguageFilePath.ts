import { existsSync } from "fs";

export default function (path: string, language: string): string {
   const languagePath = path.replace("html", language + ".html");

   if (existsSync(languagePath)) {
      return languagePath;
   }

   return path;
}

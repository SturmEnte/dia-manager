import { Client } from "pg";
import { randomBytes } from "crypto";

const tablesConfig = require("../../configs/tables.json");
const generationConfig = require("../../configs/generation.json");

// Generate a random token and check if it is already in use
// If an error occurs during the genration or checking, then "undefined" is returned else a valid token is returned
export default async function (client: Client): Promise<string | undefined> {
   let token: string;

   let start = Date.now();

   try {
      do {
         // Timeout the generation process
         if (Date.now() - start > generationConfig.password.timeoutMillis) {
            return undefined;
         }

         // Generate token
         token = randomBytes(Math.floor(tablesConfig.users.tokenLength / 2)).toString("hex");
      } while ((await client.query(`SELECT COUNT(id) FROM access_tokens WHERE token = '${token}'`)).rows[0].count > 0);
   } catch (error) {
      console.error("Error while generating access token");
      console.error(error);
      return undefined;
   }

   return token;
}

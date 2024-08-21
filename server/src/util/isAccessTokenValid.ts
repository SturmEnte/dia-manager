import { Client } from "pg";

// The function is made unsafe on purpose so that I can react to the errors in the code where I use this function
export default async function (client: Client, config: any, token: string): Promise<boolean> {
   // Query the database to find the token and its creation time
   const result = await client.query("SELECT created FROM access_tokens WHERE token = $1", [token]);

   // If no token is found, return false
   if (result.rows.length === 0) {
      return false;
   }

   // Get the creation time of the token
   const tokenCreatedAt = result.rows[0].created_at;

   // Calculate the token's age
   const tokenAge = Date.now() - new Date(tokenCreatedAt).getTime();

   // Check if the token is expired
   if (tokenAge > config.token_validity_millis) {
      return false;
   }

   return true;
}

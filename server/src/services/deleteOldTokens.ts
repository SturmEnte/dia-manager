import { Client } from "pg";

const PREFIX = "[Delete Old Tokens Service]: ";

export default async function (client: Client, config: any) {
   try {
      // Delete all tokens that are too old
      const deleteInvalidTokensQuery = await client.query("DELETE FROM access_tokens WHERE created <= NOW() - INTERVAL '1 millisecond' * $1", [config.token_validity_millis]);

      if (deleteInvalidTokensQuery.rowCount === 0) {
         console.info(PREFIX + "No invalid tokens where found and deleted");
         return;
      }
      console.info(PREFIX + "Successfully deleted " + deleteInvalidTokensQuery.rowCount + " invalid tokens");
   } catch (error) {
      console.error(PREFIX + "Error while checking for old tokens and deleting them", error);
   }
}

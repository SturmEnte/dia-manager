import { Router, json } from "express";
import { Client } from "pg";

let client: Client | undefined;

const router = Router();

router.use(json());

router.delete("/logout", async (req, res) => {
   // Check if the database connection client was successfully given when creating the router
   // If not, then the server will respond with an error
   // After that an error will be thrown that will lead to the server crashing
   // If the server is setup correctly, it will just restart the server and hopefully resolve this issue by restarting it
   if (!client) {
      res.status(500).json({ error: "Failed database connection" });
      return;
   }

   const token = req.headers.authorization;

   // Check if the username and password are both given and strings
   // The username and password should generally be strings, but this prevents unnecessary crashes
   if (!token || typeof token !== "string") {
      res.status(400).json({ error: "Username required and must be a string" });
      return;
   }

   try {
      // Try to delete the given token
      const deleteTokenQuery = await client.query("DELETE FROM access_tokens WHERE token = $1", [token]);

      // Check if there was a token that was deleted
      // If not, then the server will respond with an error
      // Otherwise, the server will respond with a success code
      if (deleteTokenQuery.rowCount == 0) {
         res.status(400).json({ error: "The token did not exist" });
         return;
      }

      res.sendStatus(200);
   } catch (error) {
      // Log the error and respond to the client with an internal server error message if an error occurres
      console.error("Error during login process", error);
      res.status(500).json({ error: "Internal server error" });
   }
});

// Get the databse connection client instance when creating the route
export default (newClient: Client): Router => {
   client = newClient;
   return router;
};

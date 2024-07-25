import { Router, json } from "express";
import { Client } from "pg";

import { hashString } from "../../../util/hashing";
import generateValidAccessToken from "../../../util/generateValidAccessToken";

let client: Client | undefined;

const router = Router();

router.use(json());

router.post("/login", async (req, res) => {
   // Check if the database connection client was successfully given when creating the router
   // If not, then the server will respond with an error
   // After that an error will be thrown that will lead to the server crashing
   // If the server is setup correctly, it will just restart the server and hopefully resolve this issue by restarting it
   if (!client) {
      res.status(500).json({ error: "Failed database connection" });
      return;
   }

   const { username, password } = req.body;

   // Check if the username and password are both given and strings
   // The username and password should generally be strings, but this prevents unnecessary crashes
   if (!username || typeof username !== "string") {
      res.status(400).json({ error: "Username required and must be a string" });
      return;
   }

   if (!password || typeof password !== "string") {
      res.status(400).json({ error: "Password required and must be a string" });
      return;
   }

   try {
      // Check if there is a user with the given username (ignoring casing)
      // If not, then return with an user error
      const userQuery = await client.query("SELECT id, password FROM users WHERE LOWER(username) = LOWER($1)", [username]);
      if (userQuery.rows.length === 0) {
         res.status(400).json({ error: "No user with that username exists" });
         return;
      }

      // Retrive the user data from the database's response
      const user = userQuery.rows[0];
      const savedHashedPassword = user.password;
      const hashedPassword = hashString(password);

      // Check if the given password and saved password match
      // If not, then an user error will be returned
      if (savedHashedPassword !== hashedPassword) {
         res.status(400).json({ error: "The password is incorrect" });
         return;
      }

      // Generate access token that is not used yet
      // If the function takes to long, it will return undefined
      // This will lead to the login process to fail and an error to be thrown
      const accessToken = await generateValidAccessToken(client);
      if (!accessToken) {
         throw new Error("Error while generating a valid access token");
      }

      // Insert the newly generated access token with the user id and creation time/date into the database
      await client.query(`INSERT INTO access_tokens (user_id, token, created) VALUES ($1, $2, NOW())`, [user.id, accessToken]);

      // Respond with the new access token, if the processes before were successfull
      res.json({ token: accessToken });
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

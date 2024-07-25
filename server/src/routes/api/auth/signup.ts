import { Router, json } from "express";
import { Client } from "pg";

import formatDate from "../../../util/formatDate";
import { hashString } from "../../../util/hashing";

const userCredentials = require("../../../../configs/usercredentials.json");

let client: Client | undefined;

const router = Router();

router.use(json());

router.post("/signup", async (req, res) => {
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

   // Check if the username and password length are within allowed limits
   // If not, return an error message indicating the problem
   if (username.length > userCredentials.maxUsernameLength) {
      res.status(400).json({ error: "The username is too long" });
      return;
   }

   if (password.length > userCredentials.maxPasswordLength) {
      res.status(400).json({ error: "The password is too long" });
      return;
   }

   const hashedPassword = hashString(password);

   try {
      // Check if there is already a user with the given username (ignoring casing)
      // If there is, return with a user error
      const userExistsQuery = await client.query("SELECT COUNT(id) FROM users WHERE LOWER(username) = LOWER($1)", [username]);
      if (userExistsQuery.rows[0].count > 0) {
         res.status(400).json({ error: "A user with that username already exists" });
         return;
      }

      // Insert the new user into the database with the hashed password and current date
      await client.query("INSERT INTO users (username, password, created) VALUES ($1, $2, $3)", [username, hashedPassword, formatDate(new Date())]);

      // Respond with a success message if the user was created successfully
      res.json({ message: "Created user" });
   } catch (error) {
      // Log the error and respond to the client with an internal server error message if an error occurs
      console.error("Error while creating user", error);
      res.status(500).json({ error: "Internal server error" });
   }
});

// Get the database connection client instance when creating the route
export default (newClient: Client): Router => {
   client = newClient;
   return router;
};

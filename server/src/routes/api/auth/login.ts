import { Router, json } from "express";
import { Client } from "pg";

import formatDate from "../../../util/formatDate";
import { hashString } from "../../../util/hashing";
import generateValidAccessToken from "../../../util/generateValidAccessToken";

let client: Client | undefined;

const router = Router();

router.use(json());

router.post("/login", async (req, res) => {
   if (!client) {
      res.status(500).json({ error: "Failed database connection" });
      return;
   }

   if (!req.body.username) {
      res.status(400).json({ error: "Username required" });
      return;
   }

   if (!req.body.password) {
      res.status(400).json({ error: "Password required" });
      return;
   }

   const username: string = req.body.username;
   const password: string = req.body.password;

   if ((await client.query(`SELECT COUNT(id) FROM users WHERE LOWER(username) = LOWER('${username}')`)).rows[0].count < 1) {
      res.status(400).json({ error: "No user with that username exists" });
      return;
   }

   let id: string;

   // Check if the entered password matches the password stored in the database
   try {
      const user = (await client.query(`SELECT id, password FROM users WHERE username = '${username}'`)).rows[0];
      id = user.id;
      const savedHashedPassword = user.password;
      const hashedPassword = hashString(password);

      if (savedHashedPassword != hashedPassword) {
         res.status(400).json({ error: "The password is incorrect" });
         return;
      }
   } catch (error) {
      if (error) {
         console.error("Error while checking user credentials");
         console.error(error);
         res.status(500).json({ error: "Error while checking user credentials" });
         return;
      }
   }

   // Create access token
   try {
      const accessToken = await generateValidAccessToken(client);

      if (!accessToken) {
         console.error("Error while generating a valid access token");
         return;
      }

      await client.query(`INSERT INTO access_tokens (user_id, token, created) VALUES (${id}, '${accessToken}','${formatDate(new Date())}')`);

      res.json({ token: accessToken });
      return;
   } catch (error) {
      if (error) {
         console.error("Error while creating access token");
         console.error(error);
         res.status(500).json({ error: "Error while creating access token" });
         return;
      }
   }
});

export default (newClient: Client): Router => {
   client = newClient;
   return router;
};

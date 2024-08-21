import "dotenv/config";
import express from "express";
import cookieParser from "cookie-parser";
import path from "path";
import pg from "pg";

import setupDatabase from "./util/setupDatabase";
import getLanguageFilePath from "./util/getLanguageFilePath";
import isAccessTokenValid from "./util/isAccessTokenValid";

import deleteOldTokens from "./services/deleteOldTokens";

import signup from "./routes/api/auth/signup";
import login from "./routes/api/auth/login";
import logout from "./routes/api/auth/logout";

const config = require("../configs/config.json");

const WHITE_LIST = ["api", "login", "signup"];
const API_WHITE_LIST = ["login", "signup"];

(async () => {
   // Connect to database
   const client = new pg.Client({
      host: process.env.PG_HOST,
      port: Number(process.env.PG_PORT),
      user: process.env.PG_USER,
      password: process.env.PG_PASSWORD,
      database: process.env.PG_DATABASE,
   });

   await client.connect((err) => {
      if (err) {
         console.error("Error while connecting to database");
         throw err;
      }

      console.log("Connected to database");
   });

   await setupDatabase(client);

   // Setup services
   setInterval(() => deleteOldTokens(client, config), config.token_validity_check_interval_millis);

   const PATH_TO_PUBLIC_FOLDER = path.join(__dirname, config.publicLocation);

   const app = express();

   app.use(express.static(PATH_TO_PUBLIC_FOLDER));

   app.use(cookieParser());

   app.all("*", (req, res, next) => {
      for (let whiteListWord of WHITE_LIST) {
         if (req.url.includes(whiteListWord)) {
            next();
            return;
         }
      }

      if (req.cookies["loggedIn"] == undefined) {
         res.redirect("/#/login");
         return;
      }

      next();
   });

   app.get("/#/*", (req, res, next) => {
      res.sendFile(path.join(PATH_TO_PUBLIC_FOLDER, "index.html"));
   });

   app.all("/api/*", async (req, res, next) => {
      // Continue to api, if the endpoint is on the whitelist
      for (let whiteListWord of WHITE_LIST) {
         if (req.url.includes(whiteListWord)) {
            next();
            return;
         }
      }

      try {
         // Check if the token is valid with the function
         if (!(await isAccessTokenValid(client, config, req.headers.authorization))) {
            // Respond with an user error if the token is invalid
            res.status(400).json({ error: "Invalid access token" });
            return;
         }
      } catch (error) {
         // Respond with an internal server error if the check fails
         console.error("Error during login process", error);
         res.status(500).json({ error: "Internal server error" });
      }

      // Continue to api if the token is valid
      next();
   });

   app.use("/api/auth/", signup(client));
   app.use("/api/auth/", login(client));
   app.use("/api/auth/", logout(client));

   app.listen(process.env.PORT, () => {
      console.log("Listening on port " + process.env.PORT);
   });
})();

import { Router, json } from "express";
import { Client } from "pg";

let client: Client | undefined;

const router = Router();

router.post("/signup", (req, res) => {});

export default (newClient: Client): Router => {
	client = newClient;
	return router;
};

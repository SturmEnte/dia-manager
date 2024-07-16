import { Router } from "express";

const router = Router();

router.post("/signup", (req: any, res) => {
	console.log(req.db);
});

export default router;

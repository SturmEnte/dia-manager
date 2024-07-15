const { Router } = require("express");

const router = Router();

router.post("/signup", (req, res) => {
	console.log(req.db);
});

module.exports = router;

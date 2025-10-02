const usernameElement = document.getElementById("username");
const passwordElement = document.getElementById("password");

document.getElementsByTagName("form")[0].onsubmit = async (e) => {
	console.log("d");
	e.preventDefault();

	const res = await fetch("/api/auth/register", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			username: usernameElement.value,
			password: passwordElement.value,
		}),
	});

	if (res.status == 200) {
		window.location.href = "/";
		return;
	}

	alert("Failed to register");
};

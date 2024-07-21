const usernameElement = document.getElementById("username");
const passwordElement = document.getElementById("password");
const errorElement = document.getElementById("error");

document.getElementById("signup-button").onclick = () => {
	fetch("/api/auth/signup", {
		method: "post",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			username: usernameElement.value,
			password: passwordElement.value,
		}),
	}).then(async (res) => {
		if (!res.headers.has("Content-Type") || !res.headers.get("Content-Type").includes("application/json")) {
			errorElement.innerText = "Error in server response";
			return;
		}

		const body = await res.json();

		if (body.error) {
			errorElement.innerText = "Error: " + body.error;
			return;
		}

		errorElement.innerText = body.message;
		errorElement.style.color = "green";
		window.location.href = "/login";
	});
};

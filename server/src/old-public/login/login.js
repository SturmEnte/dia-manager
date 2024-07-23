const usernameElement = document.getElementById("username");
const passwordElement = document.getElementById("password");
const errorElement = document.getElementById("error");

document.getElementById("login-button").onclick = () => {
	fetch("/api/auth/login", {
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

		localStorage.setItem("token", body.access_token);
		document.cookie = "loggedIn=true;path=/";

		errorElement.innerText = "Success. Redirecting shortly...";
		errorElement.style.color = "green";

		setTimeout(() => {
			window.location.href = "/";
		}, 1000);
	});
};

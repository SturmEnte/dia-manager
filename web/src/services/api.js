class DiaManagerAPIService {
	constructor(basePath = "/api") {
		this.basePath = basePath;
	}

	async request(path, init) {
		const res = await fetch(this.basePath + path, init);

		// Check if unauthorized
		if (res.status == 401) {
			// TODO: Change with cleaner info thingy
			alert("Unauthorized");
			window.location.href = "/login";
			return;
		}

		if (res.status >= 500) {
			// TODO: Change with cleaner info thingy
			alert("Server side error while requesting data");
			return;
		}

		return res;
	}

	// User
	async getUserData() {
		const res = await this.request("/user/me");

		if (!res) return;

		const data = await res.json();

		return data;
	}

	// Catheter
	async getCatheters() {
		const res = await this.request("/catheters");

		if (!res) return [];

		const data = await res.json();

		return data.catheters;
	}

	async createCatheter(start, end) {
		let catheter = { startedAt: new Date(start).toISOString() };

		if (end) catheter.endedAt = new Date(end).toISOString();

		const res = await this.request("/catheters", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(catheter),
		});

		if (!res) return;

		if (res.status === 201) {
			const id = await res.json().id;
			return { id, startedAt: start, endedAt: end };
		}
	}
}

const api = new DiaManagerAPIService();
export default api;

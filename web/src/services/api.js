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

		if (res.status >= 400 && res.status < 500) {
			// TODO: Change with cleaner info thingy
			alert("User error");
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

	async createCatheter(start, end, changeReason) {
		let catheter = { startedAt: new Date(start).toISOString() };

		if (end) catheter.endedAt = new Date(end).toISOString();
		if (changeReason) catheter.changeReason = changeReason;

		const res = await this.request("/catheters", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(catheter),
		});

		if (!res) return;

		if (res.status === 201) {
			const id = (await res.json()).id;
			return { id, startedAt: start, endedAt: end, changeReason };
		}
	}

	// id and start are required, end is optional
	async updateCatheter(id, start, end, changeReason) {
		let catheter = { startedAt: new Date(start).toISOString() };

		if (end) catheter.endedAt = new Date(end).toISOString();
		if (changeReason) catheter.changeReason = changeReason;

		const res = await this.request("/catheters/" + id, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(catheter),
		});

		if (!res) return;

		if (res.status === 204) {
			return { id, startedAt: start, endedAt: end, changeReason };
		}
	}

	async deleteCatheter(id) {
		const res = await this.request("/catheters/" + id, {
			method: "DELETE",
		});

		if (!res) return;

		if (res.status === 200) {
			return { id };
		}
	}

	// Inventory
	// Item Structures
	async getItemStructures() {
		const res = await this.request("/inventory/structures", {
			method: "get",
		});

		if (!res) return;

		if (res.status === 200) {
			let data = await res.json();

			return data;
		}
	}

	async createItemStructures(name, generallInformation, attributes) {
		const res = await this.request("/inventory/structures", {
			method: "post",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				name,
				general_information: generallInformation,
				attributes,
			}),
		});

		if (!res) return;

		if (res.status === 200) {
			let data = await res.json();

			return data.id;
		}
	}

	// Items
	async getItems() {
		const res = await this.request("/inventory/items", {
			method: "get",
		});

		if (!res) return;

		if (res.status === 200) {
			let data = await res.json();

			return data;
		}
	}
}

const api = new DiaManagerAPIService();
export default api;

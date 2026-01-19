import { createRouter, createWebHistory } from "vue-router";

import Dashboard from "./views/Dashboard.vue";
import Inventory from "./views/Inventory.vue";
import Catheters from "./views/Catheters.vue";

export const router = createRouter({
	history: createWebHistory(),
	routes: [
		{ path: "/", component: Dashboard },
		{ path: "/inventory", component: Inventory },
		{ path: "/catheters", component: Catheters },
	],
});

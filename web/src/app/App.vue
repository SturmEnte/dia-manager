<script setup>
import { ref, onMounted } from "vue";
import api from "../services/api";

const username = ref("");

onMounted(async () => {
	try {
		const userData = await api.getUserData();
		if (userData && userData.username) {
			username.value = userData.username;
		}
	} catch (error) {
		console.error("Failed to fetch user data:", error);
	}
});
</script>

<template>
	<div id="app-container">
		<div id="header">
			<div id="left-h">
				<div id="app-name">DiaManager</div>
				<nav>
					<router-link to="/">Dashboard</router-link>
					<router-link to="/catheters">Katheter</router-link>
				</nav>
			</div>
			<div>
				Welcome back <span id="username">{{ username }}</span> !
			</div>
		</div>
		<main>
			<router-view />
		</main>
	</div>
</template>

<style scoped>
#app-container {
	display: flex;
	flex-direction: column;
	height: calc(100vh - 2 * var(--padding));
}

#header {
	background: var(--col-2);
	border-radius: var(--radius);
	padding: var(--padding);
	margin-bottom: var(--padding);

	width: calc(100vw - 2 * var(--padding));

	display: flex;
	align-items: center;
	justify-content: space-between;
	flex-direction: row;
}

#header #left-h {
	display: flex;
	align-items: center;
	justify-content: flex-start;
	flex-direction: row;
}

#header #app-name {
	margin-right: var(--padding);
	font-size: 1.2rem;
}

/*router-link turns into a when rendered*/
#header nav a {
	margin-right: calc(var(--padding) / 2);
	text-decoration: none;
	color: var(--col-font);
}

#header nav a.router-link-active {
	color: var(--col-accent);
}

#username {
	color: var(--col-accent);
	font-style: italic;
}

main {
	flex: 1;
}
</style>

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
	<div>
		<div id="header">
			<nav>
				<router-link to="/">Dashboard</router-link>
				<router-link to="/catheters">Katheter</router-link>
			</nav>
			<div>{{ username }}</div>
		</div>
		<main>
			<router-view />
		</main>
	</div>
</template>

<style scoped></style>

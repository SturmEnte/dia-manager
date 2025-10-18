<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Catheter from "../components/Catheter.vue";

const cathetersObj = ref([]);

onMounted(async () => {
	let catheters = await api.getCatheters();

	// Sort by date descending
	catheters.sort((a, b) => new Date(b.startedAt) - new Date(a.startedAt));

	cathetersObj.value = catheters;
});

// Create new entry
const start = ref("");
const end = ref("");

async function createCatheter() {
	if (!start.value) return;

	const newCatheter = await api.createCatheter(start.value, end.value);

	if (!newCatheter) {
		return;
	}

	// Add new catheter to the catheters at the first position
	cathetersObj.value.unshift(newCatheter);

	// Reset form fields
	start.value = "";
	end.value = "";
}

function removeCatheter(id) {
	cathetersObj.value = cathetersObj.value.filter((c) => c.id !== id);
}
</script>

<template>
	<div id="main">
		<div id="history" class="window">
			<div class="title">Historie</div>
			<div id="catheters" class="scrollbar">
				<Catheter class="catheter" v-for="catheter in cathetersObj" :key="catheter.id" :id="catheter.id" :started-at="catheter.startedAt" :ended-at="catheter.endedAt" @deleted="removeCatheter" />
			</div>
		</div>
		<div id="create" class="window">
			<div class="title">Neuen Katheter eintragen</div>
			<form id="create-form" @submit.prevent="createCatheter">
				<label for="start">Start</label>
				<input id="start" name="start" type="datetime-local" required v-model="start" />
				<label for="end">Ende</label>
				<input id="end" name="end" type="datetime-local" v-model="end" :min="start || undefined" />
				<button type="submit">Eintragen</button>
			</form>
		</div>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;
}

#history {
	flex: 2;
	margin-right: calc(var(--padding) / 2);
	height: 100%;
	display: flex;
	flex-direction: column;
}

.title {
	font-size: 1.1rem;
	margin-bottom: var(--padding);
}

#catheters {
	flex: 1;
	overflow-y: auto;
}

.catheter {
	margin-bottom: calc(var(--padding) / 2);
}

#create {
	flex: 1;
	margin-left: calc(var(--padding) / 2);
}
</style>

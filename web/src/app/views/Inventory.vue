<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Structure from "../components/Structure.vue";

// Make inventory reactive so Vue updates the template when data arrives
const inventory = ref([]);

onMounted(async () => {
	// Map inventory
	const structures = await api.getItemStructures();

	for (let i = 0; i < structures.length; i++) {
		const structure = structures[i];

		// push objects into the reactive array and include the id on the value
		inventory.value.push({ id: structure.id, name: structure.name, attributes: structure.attributes, items: [] });
	}

	const items = await api.getItems();

	for (let i = 0; i < items.length; i++) {
		const item = items[i];

		const bucket = inventory.value.find((s) => s.id === item.structure_id);
		if (!bucket) {
			console.log("Couldn't find a structure for item with id", item.id);
			continue;
		}

		bucket.items.push(item);
	}
});
</script>

<template>
	<div id="main">
		<div id="structures">
			<Structure class="structure" v-for="structure in inventory" :key="structure.id" :id="structure.id" :name="structure.name" :items="structure.items" :attributes="structure.attributes" />
		</div>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;
}

#structures {
	display: flex;
	flex-direction: column;
	width: 70%;
}

.structure {
	margin-bottom: var(--padding);
}
</style>

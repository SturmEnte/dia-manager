<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

let inventory = new Map();

onMounted(async () => {
	// Map inventory
	const structures = await api.getItemStructures();

	for (let i = 0; i < structures.length; i++) {
		const structure = structures[i];

		inventory.set(structure.id, { name: structure.name, attributes: structure.attributes, items: [] });
	}

	const items = await api.getItems();

	for (let i = 0; i < items.length; i++) {
		const item = items[i];

		if (!inventory.has(item.structure_id)) {
			console.log("Couldn't find a structure for item with id", item.id);
			continue;
		}

		// I dont know why bucket is needed but if I write this code as a one liner it does not work
		const bucket = inventory.get(item.structure_id);
		bucket.items.push(item);
		inventory.set(item.structure_id, bucket);
	}

	console.log(inventory);
});
</script>

<template>
	<div id="main">
		<div>Test</div>
		<br />
		<p>{{ JSON.stringify(test) }}</p>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;

	color: black;
}
</style>

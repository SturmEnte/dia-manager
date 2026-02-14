<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Structure from "../components/Structure.vue";

// Make inventory reactive so Vue updates the template when data arrives
const inventory = ref([]);

const currentTab = ref("create-structure");

onMounted(async () => {
	// Map inventory
	const structures = await api.getItemStructures();

	for (let i = 0; i < structures.length; i++) {
		const structure = structures[i];

		// push objects into the reactive array and include the id on the value
		inventory.value.push({ 
			id: structure.id, 
			name: structure.name, 
			attributes: structure.attributes, 
			generalInformation: structure.general_information,
			items: [] 
		});
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

function test() {
	console.log(currentTab.value)
}
</script>

<template>
	<div id="main">
		<div id="structures" class="scrollbar">
			<Structure class="structure" v-for="structure in inventory" :key="structure.id" :id="structure.id" :name="structure.name" :items="structure.items" :attributes="structure.attributes" :generalInformation="structure.generalInformation" />
		</div>
		<div id="editor">
			<form id="selection" @input="test">
				<input type="radio" id="create-structure" value="create-structure" v-model="currentTab" />
				<label for="create-structure">Create Structure</label>

				<input type="radio" id="create-item" value="create-item" v-model="currentTab" />
				<label for="create-item">Create Item</label>
			</form>

			<div v-if="currentTab === 'create-structure'">
				Create Structure
			</div>
			
			<div v-if="currentTab === 'create-item'">
				Create Item
			</div>
		</div>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;
	overflow: hidden;
}

#editor {
	background: var(--col-2);
	flex: 1;
	margin-left: var(--padding);
	border-radius: var(--radius);
	padding: var(--padding);
}

#selection {
	widows: 100%;
	display: flex;
	flex-direction: row;
	align-items: center;
}

#selection input {
	display: none;
}

#selection label {
	background: var(--col-3);
	flex: 1;
	text-align: center;
	border-radius: var(--radius);
	margin-right: var(--padding);
	padding: var(--padding);
	user-select: none;
}

#selection label:last-of-type {
	margin: 0;
}

#selection input:checked + label {
	background: var(--col-accent);
}

#structures {
	flex: 2;
	overflow-y: auto;
}

.structure {
	margin-bottom: var(--padding);
}

.structure:last-child {
	margin: 0;
}
</style>

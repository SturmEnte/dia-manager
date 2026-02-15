<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Structure from "../components/Structure.vue";

// Make inventory reactive so Vue updates the template when data arrives
const inventory = ref([]);

// Editor
const currentTab = ref("create-structure");

// Use a unique ID generator to prevent key collisions
const generateId = () => Date.now() + Math.random();

const generallInformation = ref([]);
const attributes = ref([]);

function addGeneralInformation() {
	generallInformation.value.push({
		id: "gi" + generateId(),
		key: "",
		value: "",
	});
}

function addAttribute() {
	attributes.value.push({
		id: "att" + generateId(),
		name: "",
		required: false,
	});
}

addGeneralInformation();
addAttribute();

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
			items: [],
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
	console.log(currentTab.value);
}
</script>

<template>
	<div id="main">
		<div id="structures" class="scrollbar">
			<Structure
				class="structure"
				v-for="structure in inventory"
				:key="structure.id"
				:id="structure.id"
				:name="structure.name"
				:items="structure.items"
				:attributes="structure.attributes"
				:generalInformation="structure.generalInformation"
			/>
		</div>
		<div id="editor">
			<form id="selection" @input="test">
				<input type="radio" id="create-structure" value="create-structure" v-model="currentTab" />
				<label for="create-structure">Struktur erstellen</label>

				<input type="radio" id="create-item" value="create-item" v-model="currentTab" />
				<label for="create-item">Artikel erstellen</label>
			</form>

			<div class="editor-input scrollbar" v-if="currentTab === 'create-structure'">
				<!-- <form @submit.prevent> -->
				<label for="name">Name:</label>
				<input type="text" id="name" />

				<br /><br />

				<div>Allgemeine Informationen:</div>
				<div class="general-information-pair" v-for="(elem, index) in generallInformation" :key="elem.id">
					<input type="text" placeholder="Name/Titel" v-model="elem.key" />
					<input type="text" placeholder="Wert" v-model="elem.value" />
					<button v-if="index != 0" @click="generallInformation.splice(index, 1)">-</button>
				</div>
				<button id="add-general-information" @click="addGeneralInformation">+</button>

				<br /><br />

				<div>Attribute:</div>
				<div class="attribute" v-for="(elem, index) in attributes" :key="elem.id">
					<input type="text" placeholder="Name/Titel" v-model="elem.name" />
					<label for="required">Muss ausgefüllt werden?</label>
					<input type="checkbox" id="required" v-model="elem.required" />
					<button v-if="index != 0" @click="attributes.splice(index, 1)">-</button>
				</div>
				<button id="add-attribute" @click="addAttribute">+</button>

				<br /><br />

				<button id="create-structure">Erstellen</button>
				<!-- </form> -->
			</div>

			<div class="editor-input" v-if="currentTab === 'create-item'">Create Item</div>
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

	display: flex;
	flex-direction: column;
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
	color: var(--col-accent);
	font-weight: bold;
}

.editor-input {
	padding: var(--padding);
	overflow-y: auto;
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
